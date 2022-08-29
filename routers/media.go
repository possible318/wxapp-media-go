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
	web.Router("/wb", &controllers.MediaController{}, "Get:GetWb")
	web.Router("/demo", &controllers.MediaController{}, "Get:Demo")
	web.Router("/blog_same", &controllers.MediaController{}, "Get:BlogSame")
	web.Router("/blog_list", &controllers.MediaController{}, "Get:BlogList")
	web.Router("/recommend", &controllers.MediaController{}, "Get:Recommend")
	web.Router("/index_media", &controllers.MediaController{}, "Get:IndexMedia")

	// 七牛上传token
	web.Router("/token", &controllers.MediaController{}, "Get:QiNiuToken")
	// 相册列表
	web.Router("/photos", &controllers.MediaController{}, "GET:Photos")
	// 点赞
	web.Router("/star", &controllers.MediaController{}, "GET:Star")
	// 点踩
	web.Router("/step", &controllers.MediaController{}, "GET:Step")
	// 其他
	//web.Router("/other", &controllers.MediaController{}, "GET:Other")
	// 上传

}
