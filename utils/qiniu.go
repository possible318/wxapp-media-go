package utils

import (
	"fmt"

	"github.com/beego/beego/v2/server/web"

	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
)

var (
	accessKey, _ = web.AppConfig.String("QINIU_ACCESS_KEY") // 公钥
	secretKey, _ = web.AppConfig.String("QINIU_SECRET_KEY") // 私钥
	bucket, _    = web.AppConfig.String("QINIU_BUCKET")     // 空间
)

// GetSimpleToken 获取简单上传凭证
func GetSimpleToken() string {
	putPolicy := storage.PutPolicy{
		Scope:     bucket,
		MimeLimit: "image/*", // 只准上传图片
	}
	putPolicy.Expires = 7200 //示例2小时有效期
	mac := qbox.NewMac(accessKey, secretKey)
	upToken := putPolicy.UploadToken(mac)
	return upToken
}

// GetOverwriteToken 覆盖上传凭证
func GetOverwriteToken(keyToOverwrite string) string {
	putPolicy := storage.PutPolicy{
		Scope: fmt.Sprintf("%s:%s", bucket, keyToOverwrite),
	}
	putPolicy.Expires = 7200 //示例2小时有效期
	mac := qbox.NewMac(accessKey, secretKey)
	upToken := putPolicy.UploadToken(mac)
	return upToken
}
