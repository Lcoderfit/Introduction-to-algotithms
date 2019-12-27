package main

import (
	"Go_SendEmail/send"
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"net/mail"
	"os"
)

func main() {
	htmlBuf := HTMLContent()
	SendEmail(htmlBuf)
}

// 发送邮件
func SendEmail(buf *bytes.Buffer) {
	// 先定义一个send.Message, 然后调用NewSmtpClient获取一个Sender接口
	// 再用Sender接口调用Send或者AsyncSend发送邮件
	to := []string{"1297611505@qq.com", "lcoderfit@163.com"}
	//buf := &bytes.Buffer{}
	//fmt.Fprintf(buf, "这是邮件的内容")
	// 邮件报文的头部扩展信息，下面两个不是头部字段，所以没用
	extension := map[string]string{
		"name": "Lcoderfit",
		"age": "23",
	}

	msg := &send.Message{
		Subject: "这里邮件的标题",
		Content: buf,
		To: to,
		Extension: extension,
	}

	addr := "smtp.qq.com:25"
	from := mail.Address{
		Name: "future man",
		Address: "1297611505@qq.com",
	}
	sender, err := send.NewSmtpSender(addr, from, "mhmokifwqvagieci")
	if err != nil {
		log.Fatal(err)
	}
	err = sender.Send(msg, true)
	if err != nil {
		log.Fatal(err)
	}
}

// 返回html格式的Content
func HTMLContent() *bytes.Buffer {
	file, err := os.Open(`.\2018新年春节祝福网页.html`)
	if err != nil {
		log.Fatal("HTMLContent: ", err)
	}
	defer file.Close()
	buf := new(bytes.Buffer)
	line := bufio.NewReader(file)
	for {
		content, _, err := line.ReadLine()
		if err == io.EOF {
			break
		}
		fmt.Fprintf(buf, "%s", string(content))
	}
	return buf
}

// 单独用smtp.SendMail(addr, from, auth, to, msg)发送
// addr: 邮件服务器	from: 发送者 auth: smtp.PlainAuth返回的Auth类	to：接收者	msg：发送的数据
//func main() {
//    auth := smtp.PlainAuth("", "1297611505@qq.com", "mhmokifwqvagieci", "smtp.qq.com")
//    to := []string{"1297611505@qq.com", "lcoderfit@163.com"}
//    nickname := "这个是昵称"
//    user := "1297611505@qq.com"
//    subject := "这是邮件的标题"
//    content_type := "Content-Type: text/plain; charset=UTF-8"
//    body := "这是邮件的内容"
//    msg := []byte("To: " + strings.Join(to, ",") + "\r\nFrom: " + nickname +
//        "<" + user + ">\r\nSubject: " + subject + "\r\n" + content_type + "\r\n\r\n" + body)
//    err := smtp.SendMail("smtp.qq.com:25", auth, user, to, msg)
//    if err != nil {
//        fmt.Printf("send mail error: %v", err)
//    }
//}