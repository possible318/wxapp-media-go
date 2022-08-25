package utils

import (
	"bufio"
	"io"
	"net/http"
	"os"

	"github.com/beego/beego/v2/adapter/logs"
)

func Download(p, fileName, src string) bool {
	res, err := http.Get(src)
	if err != nil {
		logs.Error("req err", err)
		return false
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			logs.Error(err)
		}
	}(res.Body)
	// 获得get请求响应的reader对象
	reader := bufio.NewReaderSize(res.Body, 32*1024)

	file, err := os.Create(p + fileName)
	if err != nil {
		panic(err)
	}
	// 获得文件的writer对象
	writer := bufio.NewWriter(file)

	written, _ := io.Copy(writer, reader)
	logs.Info("==== download success length :%d ====", written)
	return true
}
