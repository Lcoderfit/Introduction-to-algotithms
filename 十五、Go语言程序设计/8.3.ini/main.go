package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"reflect"
	"runtime"
	"strconv"
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

type Config struct {
	MysqlConfig `ini:"mysql"`
	RedisConfig `ini:"redis"`
}

func loadIni(filename string, data interface{}) (err error) {
	// 1.data必须是指针类型
	// 且该指针所指向的元素必须为struct类型
	t := reflect.TypeOf(data)
	fmt.Println(t.Name(), t.Kind(), t.Elem(), t.Elem().Kind())
	if t.Kind() != reflect.Ptr {
		// 创建非指针err类型
		err = errors.New(fmt.Sprintf("line: %d, data param should be a pointer", GetInfo(1)))
		return
	}
	// 获取指针对应的值的类型, 必须传入一个结构体
	if t.Elem().Kind() != reflect.Struct {
		err = errors.New(fmt.Sprintf("line: %d, data param should be a struct", GetInfo(1)))
		return
	}

	// 读取文件
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}
	// 配置文件中每一行对应一个元素
	lineSlice := strings.Split(string(b), "\r\n")
	//fmt.Printf("%#v\n", lineSlice)
	var structName string
	for idx, line := range lineSlice {
		line = strings.TrimSpace(line)
		// 跳过空行
		if len(line) == 0 {
			continue
		}
		// 跳过注释行
		if strings.HasPrefix(line, ";") || strings.HasPrefix(line, "#") {
			continue
		}
		// 读取section块
		if strings.HasPrefix(line, "[") {
			// 如果不是section区，则直接返回
			if line[0] != '[' || line[len(line)-1] != ']' {
				err = fmt.Errorf("line:%d syntax error", idx+1)
				return
			}
			// 字符串去掉前后[],判断里面字符是否存在为空的可能
			sectionName := strings.TrimSpace(line[1 : len(line)-1])
			if len(sectionName) == 0 {
				err = fmt.Errorf("line %d syntax error", idx+1)
				return
			}
			// 根据sectionName去data里面找到对应的结构体
			for i := 0; i < t.Elem().NumField(); i++ {
				// panic: reflect: Field of non-struct type *main.Config
				// 记住是t.Elem().Field() 而不是t.Field(i); 不然会报非结构体类型错误
				// *main.Config说明当前的是指针，而不是main.Config结构体类型
				field := t.Elem().Field(i)
				if sectionName == field.Tag.Get("ini") {
					structName = field.Name
					fmt.Printf("找到%s对应的嵌套结构体%s\n", sectionName, structName)
				}
			}
		} else {
			// 如果不是以[开头，通过等号进行分割，左边是value右边是key
			// 当前行没有"="则报错
			if strings.Index(line, "=") == -1 || strings.HasPrefix(line, "=") {
				err = fmt.Errorf("line:%d syntax error", idx+1)
				return
			}
			index := strings.Index(line, "=")
			key := strings.TrimSpace(line[:index])
			value := strings.TrimSpace(line[index+1:])
			// 根据structName去data里面对应的嵌套结构体取出来
			v := reflect.ValueOf(data)
			// 嵌套结构体，data所指向结构体内的字段也是结构体
			structObj := v.Elem().FieldByName(structName)
			sType := structObj.Type()
			if structObj.Kind() != reflect.Struct {
				err = fmt.Errorf(fmt.Sprintf("line:%d, data中的%s字段应该是一个结构体", GetInfo(1)))
				return
			}
			// 遍历嵌套结构体的每一个字段判断，tag是否为取出的key
			var fieldName string
			var field reflect.StructField
			for i := 0; i < structObj.NumField(); i++ {
				field = sType.Field(i)
				if field.Tag.Get("ini") == key {
					// 从结构体中找到与ini配置相匹配的字段
					// 反射获取的结构体字段，均有Type，Name等多种属性
					fieldName = field.Name
					break
				}
			}

			if len(fieldName) == 0 {
				continue
			}
			// 从结构体中根据找到的字段名获取字段对象
			// 并将该字段对象设置为ini配置中的value
			fieldObj := structObj.FieldByName(fieldName)
			fmt.Println(fieldName, field.Type.Kind())
			switch field.Type.Kind() {
			case reflect.String:
				fieldObj.SetString(value)
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				var valueInt int64
				// 将字符串解析为64字节的10进制数
				valueInt, err = strconv.ParseInt(value, 10, 64)
				if err != nil {
					err = fmt.Errorf("line:%d value type error", idx+1)
					return
				}
				fieldObj.SetInt(valueInt)
			case reflect.Bool:
				var valueBool bool
				// 因为配置文件中读取的均是字符串，所以需要解析
				valueBool, err = strconv.ParseBool(value)
				if err != nil {
					err = fmt.Errorf("line:%d value type error", idx+1)
					return
				}
				fieldObj.SetBool(valueBool)
			case reflect.Float32, reflect.Float64:
				var valueFloat float64
				valueFloat, err = strconv.ParseFloat(value, 64)
				if err != nil {
					err = fmt.Errorf("line:%d value type error", idx+1)
					return
				}
				fieldObj.SetFloat(valueFloat)
			}
		}
	}
	return err
}

// n == 1, 则funcName为调用GetInfo的函数的名称，lineNo为调用GetInfo的位置
func GetInfo(n int) (lineNo int) {
	//pc, file, lineNo, ok := runtime.Caller(n)
	_, _, lineNo, ok := runtime.Caller(n)
	if !ok {
		fmt.Printf("line: %d, runtime.Caller failed", lineNo)
	}
	//funcName = runtime.FuncForPC(pc).Name()
	//fileName = path.Base(file)
	return
}

func main() {
	// reflect.ValueOf() 和 reflect.TypeOf() 的一个区别:
	//
	// v.Elem().FieldByName（或MethodByName）都只返回一个reflect.Value类型的值
	// 或者先用reflect.Value().Type()获取一个reflect.Type类型（就相当于reflect.TypeOf()返回值）
	// 然后再通过Type类型的返回值调用FieldByName或者Field等获取字段的函数，就能使用Tag了
	//
	// t.Elem().FieldByName(或MethodByName)会返回两个值，一个reflect.StructField, 另一个 ok,
	// 而reflect.StructField有一个Tag属性，可以获取Tag标签值， reflect.StructField.Tag.Get("ini")
	// 可以只传递一个未分配内存的变量, 反射也可以修改值, reflect.Value的setInt()。。。等方法可以开辟内存

	var cfg Config
	err := loadIni(
		"E:/SocialProject/Algorithms-Tags/Introduction-to-algotithms/十五、Go语言程序设计/8.3.ini/conf.ini",
		&cfg,
	)
	if err != nil {
		fmt.Printf("load ini error:%v\n", err)
		return
	}
	fmt.Printf("%#v\n", cfg)
}
