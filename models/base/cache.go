package base

import (
	"encoding/json"
	"time"
)

type CacheData struct {
	Data      interface{}
	RebuildAt int64
	ExpireAt  int64
}

func (c CacheData) IsExpired() bool {
	return time.Now().Unix() > c.ExpireAt
}

func (c CacheData) IsNeedRebuild() bool {
	return time.Now().Unix() > c.RebuildAt
}

func (c CacheData) TransData(res interface{}) {
	bytes, _ := json.Marshal(c.Data)
	_ = json.Unmarshal(bytes, res)
}

type LocalCacheData struct {
	Data      interface{}
	RebuildAt int64
	ExpireAt  int64
}

func (c LocalCacheData) IsNeedRebuild() bool {
	return time.Now().Unix() > c.RebuildAt
}

func (c LocalCacheData) IsExpired() bool {
	return time.Now().Unix() > c.ExpireAt
}
