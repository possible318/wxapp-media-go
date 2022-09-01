package controllers

import "media/services/photo"

type MediaController struct {
	BaseController
}

func (f MediaController) Recommend() {
	cxt := f.GetContext()
	svc := photo.GetBlogSvc()
	res := svc.GetRecommend(cxt)
	f.Response(0, "success", res, 0)
}

// Photos 获取图片每次10张
func (f MediaController) Photos() {
	// 翻页字段
	page, _ := f.GetInt("page", 0)
	cxt := f.GetContext()
	svc := photo.GetBlogSvc()
	res := svc.GetPhotos(cxt, page)
	f.Response(0, "success", res, 0)
}

// Like 喜欢
func (f MediaController) Like() {
	ID, _ := f.GetInt("id", 0)
	//cxt := f.GetContext()
	svc := photo.GetBlogSvc()
	res := svc.Like(ID)
	f.Response(0, "success", res, 0)
}

// Dislike 不喜欢
func (f MediaController) Dislike() {
	ID, _ := f.GetInt("id", 0)
	//cxt := f.GetContext()
	svc := photo.GetBlogSvc()
	res := svc.Dislike(ID)
	msg := "失败"
	if res {
		msg = "成功"
	}
	f.Response(0, "success", msg, 0)
}

// QiNiuToken  七牛上传token
func (f MediaController) QiNiuToken() {
	cxt := f.GetContext()
	svc := photo.GetBlogSvc()
	res := svc.QiNiuToken(cxt)
	f.Response(0, "success", res, 0)
}
