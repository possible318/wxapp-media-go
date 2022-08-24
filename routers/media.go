package routers

import (
	"media/controllers"

	"github.com/beego/beego/v2/server/web"
)

func init() {
	// 业务
	web.Router("/wb", &controllers.MediaController{}, "Get:GetWb")
	web.Router("/demo", &controllers.MediaController{}, "Get:Demo")
	web.Router("/blog_same", &controllers.MediaController{}, "Get:BlogSame")
	web.Router("/blog_list", &controllers.MediaController{}, "Get:BlogList")
	web.Router("/recommend", &controllers.MediaController{}, "Get:Recommend")
	web.Router("/index_media", &controllers.MediaController{}, "Get:IndexMedia")
	web.Router("/photos", &controllers.MediaController{}, "GET:Photos")
}
