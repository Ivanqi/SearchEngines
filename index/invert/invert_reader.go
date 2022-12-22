package invert

import (
	"./index/dict"
	"./store"
	"./mlog"
)

type InvertReader struct {
	name string
	path string

	dicReader dict.FalconStringDictReadService
	ivtReader store.FalconSearchStoreReadService
}

func NewStringInvertReader(name string, offset int64, dicStore, ivtReader store.FalconSearchStoreReadService) FalconStringInvertReadService {
	reader := &InvertReader{name: name}
	reader.dicReader = dict.NewFalconReadMap()

	if err := reader.dicReader.LoadDic(dicStore, offset); err != nil {
		mlog.Error("Load Error: %v", err)
	}

	reader.ivtReader = ivtReader

	return reader
}

func (ir *InvertReader) Fetch(key string) (FalconDocList, bool, error) {
	dv, found := ir.dicReader.Get(key)
	if !found {
		return nil, false, nil
	}

	docList := NewMemoryFalconDocList()

	by, err := ir.ivtReader.ReadFullBytes(int64(dv.Offset), int64(dv.Length * 8))
	if err != nil {
		return nil, false, err
	}

	if err := docList.FalconDecoding(by); err != nil {
		return nil, false, err
	}

	return docList, found, nil
}