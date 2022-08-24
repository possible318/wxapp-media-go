package controllers

import "media/servers/photo"

type MediaController struct {
	BaseController
}

func (f MediaController) GetWb() {
	ctx := f.GetContext()
	svc := photo.GetWbService()
	res := svc.GetWbPhoto(ctx)
	f.Response(0, "success", res, 0)
}

func (f MediaController) Demo() {
	ctx := f.GetContext()
	svc := photo.GetWbService()
	res := svc.Demo(ctx)
	f.Response(0, "success", res, 0)
}

func (f MediaController) BlogSame() {
	svc := photo.GetBlogSvc()
	res := svc.Same()
	f.Response(0, "success", res, 0)
}

func (f MediaController) BlogList() {
	ctx := f.GetContext()
	svc := photo.GetBlogSvc()
	res := svc.GetBlogList(ctx)
	f.Response(0, "success", res, 0)
}

func (f MediaController) Recommend() {
	cxt := f.GetContext()
	svc := photo.GetBlogSvc()
	res := svc.GetRecommend(cxt)
	f.Response(0, "success", res, 0)
}

func (f MediaController) IndexMedia() {
	cxt := f.GetContext()
	svc := photo.GetBlogSvc()
	res := svc.GetIndexMedia(cxt)
	f.Response(0, "success", res, 0)
}

// Photos 获取图片每次10张
func (f MediaController) Photos() {
	// 翻页字段
	page, _ := f.GetInt("page", 1)
	cxt := f.GetContext()
	svc := photo.GetBlogSvc()
	res := svc.GetPhotos(cxt, page)
	f.Response(0, "success", res, 0)
}
