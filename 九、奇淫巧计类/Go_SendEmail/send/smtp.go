package send

import (
	"bytes"
	"fmt"
	"io"
	"mime"
	"net"
	"net/mail"
	"net/smtp"
)

// SMTP邮件实例
type SmtpClient struct {
	addr string	// 邮件服务器地址
	from mail.Address	// 发送人邮箱 mail.Address{Name string, Address string}
	auth smtp.Auth	// 授权认证,smtp.PlainAuth("", from.Address, authPwd, host)
}

// 创建一个基于SMTP的邮件实例：Sender
// addr	邮件服务器地址，例如：smtp.email.qq.com:25
// from	发送者，mail.Address类型
//type Address struct {
//	Name    string // 固有名，可以为空
//	Address string // user@domain
//}
// authPwd: QQ邮箱填授权码（用于登入第三方客户端的密码），网易邮箱填密码
func NewSmtpSender(addr string, from mail.Address, authPwd string) (Sender, error) {
	smtpCli := &SmtpClient{
		addr: addr,
		from: from,
	}
	// 将host和port两部分拆开来，return （host, port, string err error）
	host, _, err := net.SplitHostPort(addr)
	if err != nil {
		return nil, err
	}
	smtpCli.auth = smtp.PlainAuth("", from.Address, authPwd, host)
	return smtpCli, nil
}


// Send发送邮件
// Message: 包含主题，内容，邮箱地址，扩展内容
func (this *SmtpClient) Send(msg *Message, isMass bool) (err error) {
	if isMass {
		err = this.massSend(msg)
	} else {
		err = this.oneSend(msg)
	}
	return
}

// AsyncSend 异步发送邮件
func (this *SmtpClient) AsyncSend(msg *Message, isMass bool, handle func(err error)) error {
	go func() {
		err := this.Send(msg, isMass)
		handle(err)
	}()
	return nil
}

// 群发邮件
func (this *SmtpClient) massSend(msg *Message) error {
	header := this.getHeader(msg.Subject)
	if msg.Extension != nil {
		for k, v := range msg.Extension {
			header[k] = v
		}
	}
	data := this.getData(header, msg.Content)
	return smtp.SendMail(this.addr, this.auth, this.from.Address, msg.To, data)
}

// oneSend 一对一顺序发送(msg.To中包含很多个接收者邮箱，逐个遍历发送消息)
func (this *SmtpClient) oneSend(msg *Message) error {
	for _, addr := range msg.To {
		header := this.getHeader(msg.Subject)
		header["To"] = addr
		if msg.Extension != nil {
			for k, v := range msg.Extension {
				header[k] = v
			}
		}
		// 获取msg.Content的数据
		data := this.getData(header, msg.Content)
		err := smtp.SendMail(this.addr, this.auth, this.from.Address, []string{addr}, data)
		if err != nil {
			return err
		}
	}
	return nil
}

// 获取邮件报文的头部, 并添加subject标题
// From Subject Mime-Version Content-Type Content-Transfer-Encoding
func (this *SmtpClient) getHeader(subject string) map[string]string {
	header := make(map[string]string)
	// 返回邮件发送者邮箱的字符串表示
	header["From"] = this.from.String()
	header["Subject"] = mime.QEncoding.Encode("utf-8", subject)
	header["Mime-Version"] = "1.0"
	header["Content-type"] = "text/html;charset=utf-8"
	// 头部编码格式，将大块数据分成多个小块进行传输
	// Quoted-Printable: 可打印字符引用编码
	header["Content-Transfer-Encoding"] = "Quoted-Printable"
	return header
}

// 获取发送的邮件报文数据
// 将邮件报文的header和Content合并，并以[]byte类型返回数据
func (this *SmtpClient) getData(header map[string]string, body io.Reader) []byte {
	buf := new(bytes.Buffer)
	for k, v := range header {
		fmt.Fprintf(buf, "%s: %s\r\n", k, v)
	}
	fmt.Fprintf(buf, "\r\n")
	io.Copy(buf, body)
	return buf.Bytes()
}