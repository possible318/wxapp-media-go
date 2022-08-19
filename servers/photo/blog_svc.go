package photo

import (
	bm "media/models/base"
	"media/models/db"
	"media/outputs"
	bs "media/servers/base"
	"media/utils"
	"sort"
	"strings"
	"time"
)

var blogSvc *BlogSvc

type BlogSvc struct {
	bs.Service
}

func GetBlogSvc() *BlogSvc {
	if blogSvc == nil {
		blogSvc = new(BlogSvc)
	}
	return blogSvc
}

func (f BlogSvc) Same() interface{} {
	model := new(db.Blog)
	itemList := make([]*db.Blog, 0)
	_, _ = model.GetQuery().All(&itemList, "ItemId", "Text", "AddTime")
	itemMap := make(map[string]string)
	var i = 0
	for _, item := range itemList {
		i++
		flag := strings.Contains(item.Text, "<a")
		if !flag {
			continue
		}
		itemMap[item.ItemID] = item.AddTime
	}
	return itemMap
}

func (f BlogSvc) GetBlogList(ctx *bm.AppContext) interface{} {
	var res interface{}
	key := "blog:img:list:group"
	var cache bm.CacheData
	if f.ExistInCache(key, ctx) {
		cache.Data = &res
		err := f.GetFromCache(key, &cache, ctx)
		if err != nil {
			res = f.buildBlogList(key, ctx)
		} else if cache.IsNeedRebuild() && utils.LockByRedis(key+":lock", 2) {
			go f.buildBlogList(key, ctx)
		}
	} else {
		res = f.buildBlogList(key, ctx)
	}
	return res
}

func (f BlogSvc) buildBlogList(key string, ctx *bm.AppContext) interface{} {
	model := new(db.Blog)
	ormList := make([]*db.Blog, 0)
	_, _ = model.GetQuery().OrderBy("-AddTime").All(&ormList)
	picMap := make(map[string][]*db.Blog)
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

func (f BlogSvc) GetRecommend(ctx *bm.AppContext) interface{} {
	var res interface{}
	key := "blog:img:list:recommend"
	var cache bm.CacheData
	if f.ExistInCache(key, ctx) {
		cache.Data = &res
		err := f.GetFromCache(key, &cache, ctx)
		if err != nil {
			res = f.buildBlogList(key, ctx)
		} else if cache.IsNeedRebuild() && utils.LockByRedis(key+":lock", 2) {
			go f.buildRecommend(key, ctx)
		}
	} else {
		res = f.buildRecommend(key, ctx)
	}
	return res
}

func (f BlogSvc) buildRecommend(key string, ctx *bm.AppContext) interface{} {
	blogList := make([]*db.Blog, 0)
	model := new(db.Blog)
	_, _ = model.GetQuery().Filter("ShowType", 1).All(&blogList)

	res := make([]*outputs.URLItem, 0)
	for _, blog := range blogList {
		item := new(outputs.URLItem)
		item.Name = blog.Pid
		item.Src = blog.Src
		res = append(res, item)
	}
	cache := new(bm.CacheData)
	cache.Data = res
	cache.ExpireAt = time.Now().Unix() + 60
	cache.RebuildAt = time.Now().Unix() + 30
	f.PutToCache(key, cache, 60, ctx)
	return res
}

func (f BlogSvc) GetIndexMedia(ctx *bm.AppContext) interface{} {
	var res interface{}
	key := "blog:img:list:index"
	var cache bm.CacheData
	if f.ExistInCache(key, ctx) {
		cache.Data = &res
		err := f.GetFromCache(key, &cache, ctx)
		if err != nil {
			res = f.buildIndexMedia(key, ctx)
		} else if cache.IsNeedRebuild() && utils.LockByRedis(key+":lock", 2) {
			go f.buildIndexMedia(key, ctx)
		}
	} else {
		res = f.buildIndexMedia(key, ctx)
	}
	return res
}

func (f BlogSvc) buildIndexMedia(key string, ctx *bm.AppContext) interface{} {
	blogList := make([]*db.Blog, 0)
	model := new(db.Blog)
	_, _ = model.GetQuery().Filter("Index", 1).All(&blogList)

	res := make([]*outputs.URLItem, 0)
	for _, blog := range blogList {
		item := new(outputs.URLItem)
		item.Name = blog.Pid
		item.Src = blog.Src
		res = append(res, item)
	}
	cache := new(bm.CacheData)
	cache.Data = res
	cache.ExpireAt = time.Now().Unix() + 60
	cache.RebuildAt = time.Now().Unix() + 30
	f.PutToCache(key, cache, 60, ctx)
	return res
}
