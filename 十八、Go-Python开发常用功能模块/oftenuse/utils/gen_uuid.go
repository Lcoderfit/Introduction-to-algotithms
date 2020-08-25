package utils

import (
	"crypto/rand"
	"encoding/base64"
	"io"
)

//生成由小写字母和数字组成的32位字符串
func GenUUID() string {
	b := make([]byte, 32)
	/*
	ReadFull从r精确地读取len(buf)字节数据填充进buf。函数返回写入的字节数和错误（如果没有读取足够的字节）。
	只有没有读取到字节时才可能返回EOF；如果读取了有但不够的字节时遇到了EOF，函数会返回ErrUnexpectedEOF。
	只有返回值err为nil时，返回值n才会等于len(buf)。
	 */
	//io.Reader是一个类型，rand.Reader源码： var Reader io.Reader
	//Reader是一个全局、共享的密码用强随机数生成器。
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return "UUID-ERROR"
	}
	//GenMD5生成固定长度的加密字符串，长度为32位
	//返回对字节流b进行base64编码后的字符串
	//URLEncoding: RFC 4648定义的另一base64编码字符集，用于URL和文件名。
	//base64: 用于编码的字符集，字符集内有64个字符
	return GenMD5(base64.URLEncoding.EncodeToString(b))
}
