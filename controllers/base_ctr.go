package controllers

import (
	"media/models/base"
	"strconv"
	"strings"

	"github.com/beego/beego/v2/server/web"
)

type BaseController struct {
	web.Controller
}

// Response 返回数据
func (f BaseController) Response(code int, msg string, data interface{}, expires int) {
	if expires > 0 {
		f.Ctx.Output.Header("Cache-Control", "public, max-age="+strconv.Itoa(expires))
	} else {
		f.Ctx.Output.Header("Cache-Control", "no-cache")
	}

	f.Data["json"] = map[string]interface{}{
		"code": code,
		"msg":  msg,
		"data": data,
	}
	err := f.ServeJSON()
	if err != nil {
		return
	}
}

// GetContext 获取公共信息
func (f BaseController) GetContext() *base.AppContext {
	ctx := new(base.AppContext)
	ctx.App = f.GetApp()
	ctx.Platform, ctx.Version = f.GetPlatformAndVersion()
	ctx.Location = ""
	ctx.Force, _ = f.GetInt("force", 0)
	ctx.Format, _ = f.GetInt("format", 0)
	ctx.Env = web.BConfig.RunMode
	ctx.AndroidChannel = f.Ctx.Input.Header("android-channel")

	return ctx
}

func (f BaseController) GetApp() string {
	app := f.GetString("app")
	if app == "" {
		app = f.Ctx.Input.Header("app")
	}
	return app
}

// GetTraceID  获取traceId
func (f BaseController) GetTraceID() string {
	traceID := f.Ctx.Input.Header("X-Trace-Id")
	if traceID == "" {
		traceID = f.Ctx.Input.Header("HTTP_X_TRACE_ID")
	}
	return traceID
}

// GetPlatformAndVersion 获取安卓渠道和版本信息
func (f BaseController) GetPlatformAndVersion() (platform string, version int) {
	platform, version = f.getPlatformAndVersion()
	return strings.ToLower(platform), version
}

// getPlatformAndVersion  获取平台和版本
func (f BaseController) getPlatformAndVersion() (platform string, version int) {
	version, _ = f.GetInt("version", 0)
	platform = f.Ctx.Input.Header("Platform")
	if platform == "" {
		platform = f.GetString("platform", "")
	}
	if version == 0 {
		version, _ = strconv.Atoi(f.Ctx.Input.Header("Version"))
	}

	if platform == "iphone" {
		platform = "ios"
	}

	return platform, version
}
