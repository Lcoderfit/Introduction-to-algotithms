package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"oftenuse/utils"
	"os"
)

// 文件名加下划线，其他不能，大写可导出，小写不可导，驼峰命名法（无论大小驼峰，从第二个单词开始都得大写）

func main() {
	// 1.生成n位随机验证码
	//genVerifyCode()

	// 2.调用阿里云api发送短信
	//smsSend()

	// 3.获取文件路径
	//genPath()

	// 4.unix时间戳转UTC
	//utils.TimeFormatExchange()

	// 5.生成图片验证码
	//genCaptcha()

	// 6.将密码进行md5加密
	//genMD5()

	// 7.将markdown内容转换为HTML
	switchMarkdownToHTML()

	// 8.Web文件上传
}

// 8.Web文件上传


func switchMarkdownToHTML() {
	content := "## This is content<br>" +
		"| | |\n" +
		"hello"
	htmlContent := utils.SwitchMarkdownToHtml(content)
	f, err := os.Create("test.html")
	defer f.Close()

	if err != nil {
		log.Fatal("Create File Error: ", err)
		return
	}
	l, err := f.WriteString(string(htmlContent))
	if err != nil {
		log.Fatal("WriteString Error: ", err)
	}
	fmt.Println(l)

	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		t, err := template.ParseFiles("test.html")
		if err != nil {
			log.Fatal("ParseFiles ERROR: ", err)
			return
		}
		t.Execute(res, nil)
	})
	fmt.Println(htmlContent)
	http.ListenAndServe(":80", nil)
	fmt.Println("end")
}

func genMD5() {
	password := "lufei123456"
	md5pwd := utils.GenMD5(password)
	fmt.Println(md5pwd)
}

func genCaptcha() {
	fmt.Println("Lcoderfit")
}

func genPath() {
	path := utils.GetPath(0)
	fmt.Println(path)
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