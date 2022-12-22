package mlog

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"
	"sync/atomic"
)

// LogLevel type
type LogLevel int32

const (
	// LevelTrace logs everything
	LevelTrace LogLevel = (1 << iota)

	// LevelInfo log Info, Warnings and Errors
	LevelInfo

	// LevelWarn logs Warning and Errors
	LevelWarn

	// LevelError logs just Errors
	LevelError
)

const MaxBytes int = 100 * 1024 * 1024
const BackupCount int = 10

type mlog struct {
	LogLevel int32

	Trace *log.Logger
	Info *log.Logger
	Warning *log.Logger
	Error *log.Logger
	Fatal *log.Logger

	LogFile *RotatingFileHandler
}

const (
	color_red = uint8(iota + 91)
	color_green
	color_yellow
	color_blue
	color_magenta	// 洋红
)

var Logger mlog

// 创建的记录器使用的DefaultFlags
var DefaultFlags = log.Ldate | log.Ltime | log.Lshortfile

// RotatingFileHandler写入日志文件，如果文件大小超过maxBytes, 它将备份当前文件并打开一个新文件
// 最大备份文件数由backupCount设置，如果备份过多，它将删除最旧的文件
type RotatingFileHandler struct {
	fd *os.File

	fileName string
	maxBytes int
	backupCount int
}

// NewRotatingFileHandler创建目录并打开日志文件
func NewRotatingFileHandler(fileName string, maxBytes int, backupCount int) (*RotatingFileHandler, error) {
	dir := path.Dir(fileName)
	os.Mkdir(dir, 0777)

	h := new(RotatingFileHandler)

	if maxBytes <= 0 {
		return nil, fmt.Error("invalid max bytes")
	}

	h.fileName = fileName
	h.maxBytes = maxBytes
	h.backupCount = backupCount

	var err error
	h.fd, err = os.OpenFile(fileName, os.O_CREATE | os.O_WRONLY | os.O_APPEND, 0666)
	if err != nil {
		return nil, err
	}

	return h, nil
}

func (h *RotatingFileHandler) Write(p []byte) (n int, err error) {
	h.doRollover()
	return h.fd.Write(p)
}

// 关闭只需关闭文件
func (h *RotatingFileHandler) Close() error {
	if h.fd != nil {
		return h.fd.Close()
	}

	return nil
}

func (h *RotatingFileHandler) doRollover() {
	f, err := h.fd.Stat()
	if err != nil {
		return
	}

	if h.maxBytes <= 0 {
		return 
	} else if f.Size() < int64(h.maxBytes) {
		return
	}

	if h.backupCount > 0 {
		h.fd.Close()

		for i := h.backupCount - 1; i > 0; i-- {
			sfn := fmt.Sprintf("%s.%d", h.fileName, i)
			dfn := fmt.Sprintf("%s.%d", h.fileName, i + 1)

			os.Rename(sfn, dfn)
		}

		dfn := fmt.Sprintf("%s.1", h.fileName)
		os.Rename(h.fileName, dfn)

		h.fd, _ = os.OpenFile(h.fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	}
}

// 开始记录
func Start(level LogLevel, path string) {
	doLogging(level, path, MaxBytes, BackupCount)
}

func StartEx(level LogLevel, path string, maxBytes, backupCount int) {
	doLogging(level, path, maxBytes, backupCount)
}

// 停止记录
func Stop() error {
	if Logger.LogFile != nil {
		return Logger.LogFile.Close()
	}

	return nil
}

// Sync将文件的当前内容提交到稳定存储
// 通常，这意味着将最近写入的数据刷新到磁盘的文件系统内存副本
func Sync() {
	if Logger.LogFile != nil {
		Logger.LogFile.fd.Sync()
	}
}

func doLogging(logLevel LogLevel, fileName string, maxBytes, backupCount int) {
	traceHandle := ioutil.Discard
	infoHandle := ioutil.Discard
	warnHandle := ioutil.Discard
	errorHandle := ioutil.Discard
	fatalHandle := ioutil.Discard

	var fileHandle *RotatingFileHandler

	switch logLevel {
	case LevelTrace:
		traceHandle = os.Stdout
		fallthrough
	case LevelInfo:
		infoHandle = os.Stdout
		fallthrough
	case LevelWarn:
		warnHandle = os.Stdout
		fallthrough
	case LevelError:
		errorHandle = os.Stderr
		fatalHandle = os.Stderr
	}

	if fileName != "" {
		var err error
		fileHandle, err = NewRotatingFileHandler(fileName, maxBytes, backupCount)
		if err != nil {
			log.Fatal("mlog: unable to create RotatingFileHandler: ", err)
		}

		if traceHandle == os.Stdout {
			traceHandle = io.MultiWriter(fileHandle, traceHandle)
		}

		if infoHandle == os.Stdout {
			infoHandle = io.MultiWriter(fileHandle, infoHandle)
		}

		if warnHandle == os.Stdout {
			warnHandle = io.MultiWriter(fileHandle, warnHandle)
		}

		if errorHandle == os.Stderr {
			errorHandle = io.MultiWriter(fileHandle, errorHandle)
		}

		if fatalHandle == os.Stderr {
			fatalHandle = io.MultiWriter(fileHandle, fatalHandle)
		}
	}

	Logger = mlog{
		Trace:   log.New(traceHandle, yellow("[TRACE]: "), DefaultFlags),
		Info:    log.New(infoHandle, green("[INFO ]: "), DefaultFlags),
		Warning: log.New(warnHandle, magenta("[WARN ]: "), DefaultFlags),
		Error:   log.New(errorHandle, red("[ERROR]: "), DefaultFlags),
		Fatal:   log.New(errorHandle, blue("[FATAL]: "), DefaultFlags),
		LogFile: fileHandle,
	}

	atomic.StoreInt32(&Logger.LogLevel, int32(logLevel))
}

// TRACE 级别
func Trace(format string, a ...interface{}) {
	Logger.Trace.Output(2, fmt.Sprintf(format, a...))
}

// INFO 级别
func Info(format string, a ...interface{}) {
	Logger.Info.Output(2, fmt.Sprintf(format, a...))
}

// WARNING 级别
func Warning(format string, a ...interface{}) {
	Logger.Warning.Output(2, magenta(fmt.Sprintf(format, a...)))
}

// ERROR 级别
func Error(format string, a ...interface{} /*err error*/) {
	Logger.Error.Output(2, red(fmt.Sprintf(format, a...)))
}

// IfError是日志的快捷函数。错误时出错
func IfError(err error) {
	if err != nil {
		Logger.Error.Output(2, fmt.Sprintf("%s\n", err))
	}
}

// FATAL 级别
func Fatal(a ...interface{}) {
	Logger.Fatal.Output(2, fmt.Sprint(a...))
	Sync()
	os.Exit(255)
}

func Fatalf(format string, a ...interface{}) {
	Logger.Fatal.Output(2, fmt.Sprintf(format, a...))
	Sync()
	os.Exit(255)
}

func FatalIfError(err error) {
	if err != nil {
		Logger.Fatal.Output(2, fmt.Sprintf("%s\n", err))
		Sync()
		os.Exit(255)
	}
}

func red(s string) string {
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", color_red, s)
}

func green(s string) string {
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", color_green, s)
}

func yellow(s string) string {
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", color_yellow, s)
}

func blue(s string) string {
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", color_blue, s)
}

func magenta(s string) string {
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", color_magenta, s)
}

func Red(s string) string {
	return red(s)
}

func Green(s string) string {
	return green(s)
}

func Yellow(s string) string {
	return yellow(s)
}

func Blue(s string) string {
	return blue(s)
}

func Magenta(s string) string {
	return magenta(s)
}
