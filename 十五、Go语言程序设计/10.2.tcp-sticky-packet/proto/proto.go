package proto

import (
	"bufio"
	"bytes"
	"encoding/binary"
)

/*
解决TCP粘包：

1.消息数据定长，例如定长100字节，不足补空格，接收方接收后每次解析100字节数据(缺点：浪费资源和带宽)
2.使用特定的分隔符区分界限：例如换行符
3.添加消息头部，用于表示消息数据的长度，接收方收到后根据消息头中的长度解析数据


大端小端：
1.大端：假设内存地址递增的方向为从左到右，那么一个数据从高位到低位字节依次按照内存地址递增的方向存放，即大端
2.小端：一个数据从低位到高位，依次按照内存地址递增的方向存放，则小端

大端：低地址存储高位字节数
小端：低地址存储低位字节数
*/

func Encode(message string) ([]byte, error) {
	// 读取消息的长度，转换为int32类型 4字节
	var length = int32(len(message))
	var pkg = new(bytes.Buffer)
	// 写入消息头
	err := binary.Write(pkg, binary.LittleEndian, length)
	if err != nil {
		return nil, err
	}

	// 写入消息实体
	err = binary.Write(pkg, binary.LittleEndian, []byte(message))
	if err != nil {
		return nil, err
	}
	return pkg.Bytes(), nil
}

func Decode(reader *bufio.Reader) (string, error) {
	// 读取消息长度
	lengthByte, _ := reader.Peek(4)
	// 使用一个字节切片作为初始化内容 创建一个Buffer缓冲区
	lengthBuf := bytes.NewBuffer(lengthByte)
	var length int32
	// 从lengthBuf中读取内容并赋给length
	err := binary.Read(lengthBuf, binary.LittleEndian, &length)
	if err != nil {
		return "", err
	}
	// Buffered返回缓冲区中现有的可读取的字节数
	if int32(reader.Buffered()) < length+4 {
		return "", err
	}

	// 读取正真的消息数据
	pack := make([]byte, int(4+length))
	_, err = reader.Read(pack)
	if err != nil {
		return "", err
	}
	// 去除头部的4个字节
	return string(pack[4:]), nil
}
