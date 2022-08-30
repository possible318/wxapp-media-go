package manage

import (
	"image"
	"media/models/db"
	"media/outputs"
	"media/services/base"
	"os"
	"strconv"
	"strings"

	"github.com/beego/beego/v2/client/orm"

	"github.com/beego/beego/v2/adapter/logs"
)

var mediaSvc *MediaSvc

type MediaSvc struct {
	base.Service
}

func GetMediaSvc() *MediaSvc {
	if mediaSvc == nil {
		mediaSvc = new(MediaSvc)
	}
	return mediaSvc
}

// Download 下载文件
func (f MediaSvc) Download() {
	//imgPath := "/Users/lulu/project/grab/photo_grab/bl/"
	model := new(db.Blog)
	ormList := make([]*db.Blog, 0)
	_, _ = model.GetQuery().Filter("Id__gt", 3508).OrderBy("-AddTime").All(&ormList)
	for _, item := range ormList {
		logs.Info("==== download add %s ====", strconv.Itoa(item.ID))
		//pid := item.Pid
		//src := item.Src
		//utils.Download(imgPath, pid, src)
	}
}

// Upload 上传文件
func (f MediaSvc) Upload() {

}

// GetSame 获取相似
func (f MediaSvc) GetSame() interface{} {
	list := new(db.Same).GetSameMsg()
	return list
}

func (f MediaSvc) ImgMsg() interface{} {
	path := "/Users/lulu/project/grab/photo_grab/bl/"
	files, _ := os.ReadDir(path)
	a := make([]string, 0)
	for _, filepath := range files {
		if filepath.Name() == ".DS_Store" {
			continue
		}
		file, _ := os.Open(path + filepath.Name())

		img, _, err := image.DecodeConfig(file)
		if err != nil {
			logs.Error(err)
		}
		if img.Width == 0 {
			logs.Info(filepath.Name())
			a = append(a, filepath.Name())
		}
		blogOrm := new(db.Blog)
		_, _ = blogOrm.GetQuery().Filter("Pid", filepath.Name()).Update(orm.Params{"Height": img.Height, "Width": img.Width})
		err = file.Close()
		if err != nil {
			return nil
		}
	}
	ss := strings.Join(a, ",")
	print(ss)
	return nil
}

func (f MediaSvc) GetImgByItemID(itemID string) interface{} {
	blogOrm := new(db.Blog)
	blogList := make([]*db.Blog, 0)
	_, _ = blogOrm.GetQuery().Filter("ItemId", itemID).All(&blogList)

	itemList := make([]outputs.URLItem, 0)
	for _, blog := range blogList {
		item := outputs.URLItem{
			Pid: blog.Pid,
			Src: blog.Src,
		}
		itemList = append(itemList, item)
	}
	return itemList
}

func (f MediaSvc) GetSmallImg() interface{} {
	blogOrm := new(db.Blog)
	blogList := make([]*db.Blog, 0)
	_, _ = blogOrm.GetQuery().Filter("Width__lt", 800).OrderBy("Width").All(&blogList)

	itemList := make([]outputs.URLItem, 0)
	for _, blog := range blogList {
		item := outputs.URLItem{
			ID:     blog.ID,
			ItemID: blog.ItemID,
			Pid:    blog.Pid,
			Src:    blog.Src,
		}
		itemList = append(itemList, item)
	}
	return itemList
}
