package base

// AppContext 公共信息
type AppContext struct {
	TraceID        string // 请求追踪ID
	App            string // 请求来源app
	Platform       string // 平台: Android, iOS
	Version        int    // 版本信息
	IsReview       bool   // 是否审核版本
	Lang           string // 语言
	Region         string // 地区信息
	Location       string // 地理位置
	Force          int    // 强刷
	Format         int    // 格式
	Env            string // 运行环境
	DeviceID       string // 设备ID
	Referer        string // 上次路径
	Reform         string // 格式referer
	AndroidChannel string // 安卓渠道
}
