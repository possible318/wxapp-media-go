package photo

import (
	bm "media/models/base"
	"media/models/db"
	"media/outputs"
	bs "media/servers/base"
	"media/utils"
	"sort"
	"strconv"
	"strings"
	"time"
)

var wbService *WbService

type WbService struct {
	bs.Service
}

func GetWbService() *WbService {
	if wbService == nil {
		wbService = new(WbService)
	}
	return wbService
}

func (f WbService) GetWbPhoto(ctx *bm.AppContext) interface{} {
	var res interface{}
	key := "photo::list:" + strconv.Itoa(ctx.Format) + ":" + ctx.App + ":" + ctx.Lang + ":" + ctx.Env
	var cache bm.CacheData
	if f.ExistInCache(key, ctx) {
		cache.Data = &res
		err := f.GetFromCache(key, &cache, ctx)
		if err != nil {
			res = f.buildWbPhoto(key, ctx)
		} else if cache.IsNeedRebuild() && utils.LockByRedis(key+":lock", 2) {
			go f.buildWbPhoto(key, ctx)
		}
	} else {
		res = f.buildWbPhoto(key, ctx)
	}
	return res
}

func (f WbService) buildWbPhoto(key string, ctx *bm.AppContext) interface{} {
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
			urlItem.Name = i.Pid
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

func (f WbService) GetWbPhotoByTime(ctx *bm.AppContext) interface{} {
	key := ""
	return f.buildWbPhotoByTime(key, ctx)
}

func (f WbService) buildWbPhotoByTime(key string, ctx *bm.AppContext) interface{} {
	return nil
}

func (f WbService) Demo(ctx *bm.AppContext) interface{} {
	model := new(db.WbPhoto)
	itemList := make([]*db.WbPhoto, 0)
	_, _ = model.GetQuery().All(&itemList, "ItemID", "Text")
	itemMap := make(map[string]string)
	for _, item := range itemList {
		flag := strings.Contains(item.Text, "<a")
		if !flag {
			continue
		}
		itemMap[item.ItemID] = item.Text
		// itemId := strings.Split(item.(string), "_-_")
		// _, _ = model.GetQuery().Filter("ItemID", item).Update(orm.Params{"item_id": itemId[1]})
	}
	return itemMap
}
