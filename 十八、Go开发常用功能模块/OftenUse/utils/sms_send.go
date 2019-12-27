package utils

import (
	"encoding/json"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
	"log"
)

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
	return nil
}
