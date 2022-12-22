package dict

import (
	"testing"
	"./store"
	"./mlog"
	"./message"
	"fmt"
)

func Test_DictMap(t *testing.T) {
	mlog.Start(mlog.LevelInfo, "dictmap.log")

	mapWrite := NewFalconWriteMap()
	mapWrite.Put("hello", &message.DictValue{Offset:100, Length: 200})
	mapWrite.Put("hello2", &message.DictValue{Offset:200, Length:202})
	mapWrite.Put("hello3", &message.DictValue{Offset:300, Length:203})

	for i := 0; i < 1000; i++ {
		mapWrite.Put(fmt.Sprintf("hello%d", i), &message.DictValue{Offset:100, Length: 200})
	}

	mlog.Info("Write to file...")
	fmapWriter := store.NewFalconFileStoreWriteService("./map.dict")
	mapWrite.Persistence(fmapWriter)
	fmapWriter.Close()

	fmapWriter := store.NewFalconSearchStoreReadService("./map.dict")
	mapReader := NewFalconReadMap()

	if err := mapReader.LoadDic(fmapReader, 0); err != nil {
		mlog.Error("load error %v", err)
		return
	}

	fmapReader.Close()
	mlog.Info("%s", mapReader.ToString())
}