package routers

import (
	"media/controllers"

	"github.com/beego/beego/v2/server/web"
)

func init() {
	// 管理
	web.Router("/manage/media/download", &controllers.ManageController{}, "Get:Download")
	web.Router("/manage/media/same", &controllers.ManageController{}, "Get:GetSame")
	web.Router("/manage/media/img_msg", &controllers.ManageController{}, "Get:ImgMsg")
	web.Router("/manage/media/img_by_item", &controllers.ManageController{}, "Get:GetImgByItemID")
	web.Router("/manage/media/small_img", &controllers.ManageController{}, "Get:GetSmallImg")

	// 业务
	web.Router("/recommend", &controllers.MediaController{}, "Get:Recommend")
	// 相册列表
	web.Router("/photos", &controllers.MediaController{}, "GET:Photos")
	// 点赞
	web.Router("/like", &controllers.MediaController{}, "GET:Like")
	// 点踩
	web.Router("/dislike", &controllers.MediaController{}, "GET:Dislike")
	// 七牛上传token
	web.Router("/token", &controllers.MediaController{}, "Get:QiNiuToken")

}
