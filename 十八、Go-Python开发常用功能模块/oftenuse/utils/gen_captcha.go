//package utils
//
//import (
//	"encoding/json"
//	"github.com/mojocn/base64Captcha"
//	"log"
//	"net/http"
//)
//
////configJsonBody json request body.
//type configJsonBody struct {
//	Id              string
//	CaptchaType     string
//	VerifyValue     string
//	ConfigAudio     base64Captcha.ConfigAudio
//	ConfigCharacter base64Captcha.ConfigCharacter
//	ConfigDigit     base64Captcha.ConfigDigit
//}
//
//// base64Captcha create http handler
//func GenerateCaptchaHandler(w http.ResponseWriter, r *http.Request) {
//	//parse request parameters
//	//接收客户端发送来的请求参数
//	decoder := json.NewDecoder(r.Body)
//	var postParameters configJsonBody
//	err := decoder.Decode(&postParameters)
//	if err != nil {
//		log.Println(err)
//	}
//	defer r.Body.Close()
//
//	//create base64 encoding captcha
//	//创建base64图像验证码
//
//	var config interface{}
//	switch postParameters.CaptchaType {
//	case "audio":
//		config = postParameters.ConfigAudio
//	case "character":
//		config = postParameters.ConfigCharacter
//	default:
//		config = postParameters.ConfigDigit
//	}
//	captchaId, captcaInterfaceInstance := base64Captcha.GenerateCaptcha(postParameters.Id, config)
//	base64blob := base64Captcha.CaptchaWriteToBase64Encoding(captcaInterfaceInstance)
//
//	//or you can just write the captcha content to the httpResponseWriter.
//	//before you put the captchaId into the response COOKIE.
//	//captcaInterfaceInstance.WriteTo(w)
//
//	//set json response
//	//设置json响应
//	w.Header().Set("Content-Type", "application/json; charset=utf-8")
//	body := map[string]interface{}{"code": 1, "data": base64blob, "captchaId": captchaId, "msg": "success"}
//	json.NewEncoder(w).Encode(body)
//}
//
//// base64Captcha verify http handler
//func CaptchaVerifyHandle(w http.ResponseWriter, r *http.Request) {
//
//	//parse request parameters
//	//接收客户端发送来的请求参数
//	decoder := json.NewDecoder(r.Body)
//	var postParameters configJsonBody
//	err := decoder.Decode(&postParameters)
//	if err != nil {
//		log.Println(err)
//	}
//	defer r.Body.Close()
//	//verify the captcha
//	//比较图像验证码
//	verifyResult := base64Captcha.VerifyCaptcha(postParameters.Id, postParameters.VerifyValue)
//
//	//set json response
//	//设置json响应
//	w.Header().Set("Content-Type", "application/json; charset=utf-8")
//	body := map[string]interface{}{"code": "error", "data": "验证失败", "msg": "captcha failed"}
//	if verifyResult {
//		body = map[string]interface{}{"code": "success", "data": "验证通过", "msg": "captcha verified"}
//	}
//	json.NewEncoder(w).Encode(body)
//}
package utils

import "fmt"

func Test() {
	fmt.Println("Lcoderfit")
}
