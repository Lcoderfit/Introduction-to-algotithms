package utils

import (
	"fmt"
	"github.com/astaxie/beego"
	"io"
	"os"
	"path/filepath"
	"time"
)

type UploadController struct {
	// 此处可以改成BaseController
	// BaseController继承beego.Controller
	beego.Controller
}

func (this *UploadController) Post() {
	fileData, fileHeader, err := this.GetFile("upload")
	if err != nil {
		beego.Error(err)
		return
	}
	// 文件名和大小
	beego.Info("name: ", fileHeader.Filename, fileHeader.Size)

	now := time.Now()
	// 打印文件名称的后缀
	beego.Info("ext: ", filepath.Ext(fileHeader.Filename))
	fileType := "other"
	// 判断后缀为图片的文件，如果是图片格式才存入数据库中
	fileExt := filepath.Ext(fileHeader.Filename)
	if fileExt == ".jpg" || fileExt == ".png" || fileExt == ".gif" || fileExt == ".jpeg" {
		fileType = "img"
	}
	// 文件夹路径
	// 将文件按天存放
	fileDir := fmt.Sprintf("static/upload/%s/%d/%d/%d", fileType, now.Year(), now.Month(), now.Day())
	// ModePerm是0777，这样拥有该文件夹的执行权限
	// 创建文件夹并设置权限
	err = os.MkdirAll(fileDir, os.ModePerm)
	if err != nil {
		beego.Error(err)
		return
	}

	// 文件路径
	timeStamp := time.Now().Unix()
	fileName := fmt.Sprintf("%d-%s", timeStamp, fileHeader.Filename)
	filePathStr := filepath.Join(fileDir, fileName)
	// 创建文件
	desFile, err := os.Create(filePathStr)
	if err != nil {
		beego.Error(err)
		return
	}

	// 将浏览器客户端上传的文件拷贝到本地路径的文件里面
	_, err = io.Copy(desFile, fileData)
	if err != nil {
		beego.Error(err)
		return
	}

	this.Data["json"] = map[string]interface{}{
		"code": 1,
		"message": "上传成功",
	}
	this.ServeJSON()
}