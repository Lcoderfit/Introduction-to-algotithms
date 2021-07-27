package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"reflect"
	"strings"
)

type MysqlConfig struct {
	Address  string `ini:"address"`
	Port     int    `ini:"port"`
	Username string `ini:"username"`
	Password string `ini:"password"`
}

type RedisConfig struct {
	Host     string `ini:"host"`
	Port     int    `ini:"port"`
	Password string `ini:"password"`
	Database int    `ini:"database"`
}

func loadIni(filename string, data interface{}) (err error) {
	// 1.
	t := reflect.TypeOf(data)
	fmt.Println(t.Name(), t.Kind())
	if t.Kind() != reflect.Ptr {
		fmt.Errorf("%s", "必须传入指针类型")
		return
	}
	// 获取指针对应的值的类型, 必须传入一个结构体
	if t.Elem().Kind() != reflect.Struct {
		err = errors.New("data param should be a struct")
		return
	}

	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}
	// 每一个元素对应一行
	lineSlice := strings.Split(string(b), "\n")
	fmt.Printf("%#v\n", lineSlice)

	for idx, line := range lineSlice {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, ";") || strings.HasPrefix("#") {
			continue
		}
	}

	return
}

func main() {
	var mc MysqlConfig
	err := loadIni("./conf.ini", &mc)
	if err != nil {
		fmt.Println("load ini error")
	}
	fmt.Printf(mc.Password, mc.Address)
}
