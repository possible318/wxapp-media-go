package routers

import (
	"net/http"

	"github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/filter/cors"
)

func init() {
	// 跨域
	var allowOrigin = []string{"*"}

	var allowMethod = []string{"GET", "POST", "OPTIONS"}

	var allowHeader = []string{
		"Authorization", "Content-Type", "Content-Length", "Transfer-Encoding",
		"Accept", "Origin", "User-Agent", "DNT", "Cache-Control", "X-Mx-ReqToken",
		"X-Requested-With"}

	web.InsertFilter("*", web.BeforeRouter, cors.Allow(&cors.Options{
		AllowOrigins:     allowOrigin,
		AllowMethods:     allowMethod,
		AllowHeaders:     allowHeader,
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		AllowCredentials: true,
	}))

	web.ErrorHandler("404", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	})
}
