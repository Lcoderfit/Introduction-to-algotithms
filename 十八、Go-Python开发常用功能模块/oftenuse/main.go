package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"oftenuse/utils"
	"os"
	"runtime"
)

//文件名加下划线，其他不能，大写可导出，小写不可导，驼峰命名法（无论大小驼峰，从第二个单词开始都得大写）

func main() {
	//1.生成n位随机验证码
	//genVerifyCode()

	//2.调用阿里云api发送短信
	//smsSend()

	//3.获取文件路径
	//genPath()

	//4.unix时间戳转UTC
	//utils.TimeFormatExchange()

	//5.生成图片验证码
	//genCaptcha()

	//6.将密码进行md5加密
	//genMD5()

	//7.将markdown内容转换为HTML
	//switchMarkdownToHTML()

	//8.Web文件上传

	//9.验证手机号
	//checkPhone()

	//10.验证密码
	//checkPassword()

	//11.判断字符串中是否有中文
	//hasChineseChar()

	//12.Max方法
	//max()


	//13,生成UUID
	//genUUID()

	//14.net/http的r和w的设置cookie和session的操作

	//15.获取函数信息
	getFuncInfo()



	//1000
	//test1("aksdjf", " ajskdfj")
}

func test1(val ...interface{}) {
	a := fmt.Sprint(val...)
	fmt.Println(a)
}

//15.获取函数信息
func getFuncInfo() {
	utils.GetLevelCallerFuncName()

	pc, _, _, _ := runtime.Caller(0)
	fn := runtime.FuncForPC(pc).Name()
	ret := utils.GetFuncName(fn, ".")
	fmt.Println("分割后的函数名：", ret)
}

//14.net/http的r和w的设置cookie和session的操作
func setAndGetCookie() {

}

//13,生成UUID
func genUUID() {
	ret := utils.GenUUID()
	fmt.Println("GenUUID: ", ret)
}

//12.Max方法
func max() {
	var a, b int
	fmt.Scan(&a, &b)
	fmt.Println("The max number is: ", utils.Max(a, b))
}

//11.判断字符串中是否有中文
func hasChineseChar() {
	var str string
	for {
		fmt.Scanf("%s\n", &str)
		fmt.Println(utils.HasChineseChar(str))
	}
}

//10.验证密码
func checkPassword() {
	var password string
	for {
		fmt.Scanf("%s\n", &password)
		if password == "-1" {
			break
		}
		ok, err := utils.CheckPassword(password)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println(ok)
		}
	}
}

//9.验证手机号
func checkPhone() {
	fmt.Println("CheckPhone Begin")
	var phone string
	var n int

	for {
		//输入的正确方法，后面需要加一个\n, 如果不加\n会读取两次，
		//第一次读取phone, 第二次读取输入时按下回车键后末尾的\n,
		//所以同一个输出语句会打印两次
		n, _ = fmt.Scanf("%s\n", &phone)
		fmt.Println("n: ", n)
		res := utils.CheckPhone(phone)
		fmt.Printf("phone=%s, is %t\n", phone, res)
	}
	//phone := "1287049"
	//res := utils.CheckPhone(phone)
	//fmt.Println("res: ", res)
}

//8.Web文件上传


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
	//短信签名
	signName := ""
	//模板代码
	templateCode := ""
	//手机号码
	phoneNumber := ""

	err := utils.SendSms(regionId, appKey, appSecret, signName, templateCode, phoneNumber)
	if err != nil {
		log.Fatal(err)
	}
}