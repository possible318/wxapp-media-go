package photo

import (
	bm "media/models/base"
	"media/models/db"
	"media/outputs"
	bs "media/services/base"
	"media/utils"
	"sort"
	"strconv"
	"time"
)

var blService *BlService

type BlService struct {
	bs.Service
}

func GetBlService() *BlService {
	if blService == nil {
		blService = new(BlService)
	}
	return blService
}

func (f BlService) GetWbPhoto(ctx *bm.AppContext) interface{} {
	var res interface{}
	key := "photo:bl:list:" + strconv.Itoa(ctx.Format) + ":" + ctx.App + ":" + ctx.Lang + ":" + ctx.Env
	var cache bm.CacheData
	if f.ExistInCache(key, ctx) {
		cache.Data = &res
		err := f.GetFromCache(key, &cache, ctx)
		if err != nil {
			res = f.buildBlPhoto(key, ctx)
		} else if cache.IsNeedRebuild() && utils.LockByRedis(key+":lock", 2) {
			go f.buildBlPhoto(key, ctx)
		}
	} else {
		res = f.buildBlPhoto(key, ctx)
	}
	return res
}

func (f BlService) buildBlPhoto(key string, ctx *bm.AppContext) interface{} {
	orm := new(db.WbPhoto)
	ormList := make([]*db.WbPhoto, 0)
	_, _ = orm.GetQuery().OrderBy("-AddTime").All(&ormList)
	picMap := make(map[string][]*db.WbPhoto)
	for _, item := range ormList {
		picMap[item.ItemID] = append(picMap[item.ItemID], item)
	}

	itemList := make([]outputs.PicItem, 0)
	for k, v := range picMap {
		item := new(outputs.PicItem)
		item.ID = k
		picList := make([]outputs.URLItem, 0)
		for _, i := range v {
			item.Text = i.Text
			item.AddTime = i.AddTime
			urlItem := new(outputs.URLItem)
			urlItem.Pid = i.Pid
			urlItem.Text = i.Text
			urlItem.Src = i.Src
			picList = append(picList, *urlItem)
		}
		item.SrcList = picList
		itemList = append(itemList, *item)
	}
	sort.Sort(outputs.PicItemList(itemList))

	cache := new(bm.CacheData)
	cache.Data = itemList
	cache.ExpireAt = time.Now().Unix() + 60
	cache.RebuildAt = time.Now().Unix() + 30
	f.PutToCache(key, cache, 60, ctx)
	return itemList
}
