package photo

import (
	"math"
	bm "media/models/base"
	"media/models/db"
	bs "media/services/base"
	"media/types"
	"media/utils"
	"strconv"
	"time"

	"github.com/beego/beego/v2/adapter/logs"
	"github.com/beego/beego/v2/client/orm"
)

//import "github.com/qiniu/go-sdk/v7"

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

// GetPhotos 获取相册
func (f BlogSvc) GetPhotos(ctx *bm.AppContext, page int) interface{} {
	var res interface{}
	key := "blog:img:photos:page:" + strconv.Itoa(page)
	var cache bm.CacheData
	if f.ExistInCache(key, ctx) {
		cache.Data = &res
		err := f.GetFromCache(key, &cache, ctx)
		if err != nil {
			res = f.buildPhotos(page, key, ctx)
		} else if cache.IsNeedRebuild() && utils.LockByRedis(key+":lock", 2) {
			go f.buildPhotos(page, key, ctx)
		}
	} else {
		res = f.buildPhotos(page, key, ctx)
	}
	return res
}

// buildPhotos 构建相册数据
func (f BlogSvc) buildPhotos(page int, key string, ctx *bm.AppContext) *types.PhotoOpt {
	// 每页10张
	offset := page * 10
	// 获取数据
	blogList := make([]*db.Blog, 0)
	model := new(db.Blog)
	count, _ := model.GetQuery().Count()
	_, _ = model.GetQuery().OrderBy("-ID").Limit(10, offset).All(&blogList)

	srcList := make([]types.URLItem, 0)
	for _, blog := range blogList {
		item := new(types.URLItem)
		item.ID = blog.ID
		item.Pid = blog.Pid
		item.ItemID = blog.ItemID
		item.Text = blog.Text
		item.Src = blog.Src
		item.Like = blog.Like
		item.Dislike = blog.Dislike
		srcList = append(srcList, *item)
	}

	res := new(types.PhotoOpt)
	res.Page = page
	res.Total = int(math.Ceil(float64(count) / 10.0))
	res.List = srcList

	cache := new(bm.CacheData)
	cache.Data = res
	cache.ExpireAt = time.Now().Unix() + 60
	cache.RebuildAt = time.Now().Unix() + 30
	f.PutToCache(key, cache, 60, ctx)
	return res
}

// GetRecommend 获取推荐
func (f BlogSvc) GetRecommend(ctx *bm.AppContext) interface{} {
	var res interface{}
	key := "blog:img:list:recommend"
	var cache bm.CacheData
	if f.ExistInCache(key, ctx) {
		cache.Data = &res
		err := f.GetFromCache(key, &cache, ctx)
		if err != nil {
			res = f.buildRecommend(key, ctx)
		} else if cache.IsNeedRebuild() && utils.LockByRedis(key+":lock", 2) {
			go f.buildRecommend(key, ctx)
		}
	} else {
		res = f.buildRecommend(key, ctx)
	}
	return res
}

// buildRecommend 构建推荐数据
func (f BlogSvc) buildRecommend(key string, ctx *bm.AppContext) interface{} {
	blogList := make([]*db.Blog, 0)
	model := new(db.Blog)
	_, _ = model.GetQuery().Filter("ShowType", 1).All(&blogList)

	res := make([]*types.URLItem, 0)
	for _, blog := range blogList {
		item := new(types.URLItem)
		item.Pid = blog.Pid
		item.Text = blog.Text
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

// Like  点赞
func (f BlogSvc) Like(id int) bool {
	model := new(db.Blog)
	info, err := model.GetById(id)
	if err != nil {
		logs.Error(err)
		return false
	}
	info.Like += 1
	_, err = model.GetQuery().Filter("Id", id).Update(orm.Params{"like": info.Like})
	if err != nil {
		logs.Error(err)
		return false
	}
	return true
}

// Dislike Step 点踩
func (f BlogSvc) Dislike(id int) bool {
	model := new(db.Blog)
	info, err := model.GetById(id)
	if err != nil {
		logs.Error(err)
		return false
	}
	info.Dislike += 1
	_, err = model.GetQuery().Filter("Id", id).Update(orm.Params{"dislike": info.Dislike})
	if err != nil {
		logs.Error(err)
		return false
	}
	return true
}

// QiNiuToken 上传图片到七牛云
func (f BlogSvc) QiNiuToken(ctx *bm.AppContext) string {
	var res string
	key := "Blog:Media:QiNiuToken"
	var cache bm.CacheData
	if f.ExistInCache(key, ctx) {
		cache.Data = &res
		err := f.GetFromCache(key, &cache, ctx)
		if err != nil {
			res = utils.GetSimpleToken()
			cache := new(bm.CacheData)
			cache.Data = res
			cache.ExpireAt = time.Now().Unix() + 3700
			cache.RebuildAt = time.Now().Unix() + 3600
			f.PutToCache(key, cache, 3600, ctx)
		}
	} else {
		res = utils.GetSimpleToken()
		cache := new(bm.CacheData)
		cache.Data = res
		cache.ExpireAt = time.Now().Unix() + 3700
		cache.RebuildAt = time.Now().Unix() + 3600
		f.PutToCache(key, cache, 3600, ctx)
	}
	return res
}
