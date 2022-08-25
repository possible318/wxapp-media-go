package controllers

import (
	"media/servers/manage"
)

type ManageController struct {
	BaseController
}

func (f ManageController) Download() {
	svc := manage.GetMediaSvc()
	svc.Download()
}

func (f ManageController) GetSame() {
	svc := manage.GetMediaSvc()
	res := svc.GetSame()
	f.Response(0, "success", res, 0)
}

func (f ManageController) ImgMsg() {
	svc := manage.GetMediaSvc()
	res := svc.ImgMsg()
	f.Response(0, "success", res, 0)
}

func (f ManageController) GetImgByItemID() {
	itemID := f.GetString("item_id")
	svc := manage.GetMediaSvc()
	res := svc.GetImgByItemID(itemID)
	f.Response(0, "success", res, 0)
}

func (f ManageController) GetSmallImg() {
	svc := manage.GetMediaSvc()
	res := svc.GetSmallImg()
	f.Response(0, "success", res, 0)
}
