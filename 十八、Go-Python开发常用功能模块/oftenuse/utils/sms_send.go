package utils

import (
	"encoding/json"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
	"log"
)

// 可以把返回值改为code string，然后if err != nil, return ""
func SendSms(regionId, appKey, appSecret, signName, templateCode, phoneNumber string) error {
	client, err := dysmsapi.NewClientWithAccessKey(regionId, appKey, appSecret)
	if err != nil {
		log.Printf("NewClientWithAccessKey Error: ")
		return err
	}

	request := dysmsapi.CreateSendSmsRequest()
	request.Scheme = "https"

	request.SignName = signName
	request.TemplateCode = templateCode
	request.PhoneNumbers = phoneNumber

	// 生成6位验证码
	code := GenVerifyCode(6)
	par, err := json.Marshal(map[string]interface{}{
		"code": code,
	})
	request.TemplateParam = string(par)

	response, err := client.SendSms(request)
	if err != nil {
		log.Printf("SendSms Error: ")
		return err
	}
	log.Println(response)

	// 如果返回值改成 (code string), 添加如下代码
	// response.Code表示请求的状态，请求成功返回"OK"
	//if response.Code != "OK" {
	//	return ""
	//}

	return nil
}
