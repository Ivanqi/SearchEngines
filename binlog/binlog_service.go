package binlog

import (
	"./store"
	"sync"
)

type FalconBinlogService interface {
	AppendLog(logMessage *message.BinlogMessage) (int64, error)
}

type FalconBinlog struct {
	storeWriteService store.FalconSearchStoreWriteService
	longId int64
	locker *sync.RWMutex
}

type NewFalconBinlog() FalconBinlogService {
	fd := &FalconBinlog{logId:0, locker:new(sync.RWMutex)}
	fd.storeWriteService = store.NewFalconFileStoreWriteService("./bin.log")
	return fb
}

func (fb *FalconBinlog) AppendLog(logMessage *message.BinlogMessage) (int64, error) {
	fb.locker.Lock()
	defer fb.locker.Unlock()

	logMessage.LogId = fb.logId
	fb.logId++
	fb.storeWriteService.AddFieldMapping(logMessage)
	return fb.logId, nil
}