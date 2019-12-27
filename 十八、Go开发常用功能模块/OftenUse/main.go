package main

import (
	"fmt"
	"github.com/astaxie/beego/plugins/apiauth"
	"github.com/gpmgo/gopm/modules/log"
	"oftenuse/utils"
)

// 文件名加下划线，其他不能，大写可导出，小写不可导，驼峰命名法（无论大小驼峰，从第二个单词开始都得大写）

func main() {
	// 生成n位随机验证码
	//genVerifyCode()

	//调用阿里云api发送短信
	//smsSend()
}

func genVerifyCode() {
	code := utils.GenVerifyCode(6)
	fmt.Println("code:", code)
}

func smsSend() {
	regionId := "cn-hangzhou"
	appKey := ""
	appSecret := ""
	// 短信签名
	signName := ""
	// 模板代码
	templateCode := ""
	// 手机号码
	phoneNumber := ""

	err := utils.SendSms(regionId, appKey, appSecret, signName, templateCode, phoneNumber)
	if err != nil {
		log.Fatal(err)
	}
}