package utils

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"io"
)

// 密码加密
func GenMD5(pwd string) string {
	md5password := md5.New()
	io.WriteString(md5password, pwd)
	// 初始化一个Buffer，返回*Buffer类型
	// Buffer：缓冲区 的作用：减少IO次数，数据缓冲到一定量再一次性读
	buffer := bytes.NewBuffer(nil)
	fmt.Fprintf(buffer, "%x", md5password.Sum(nil))
	// 将Buffer中的未读取数据以字符串形式返回
	return buffer.String()
}