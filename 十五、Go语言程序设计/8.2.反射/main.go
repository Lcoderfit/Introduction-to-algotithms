package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type person struct {
	Name string
	Age  int `json:"age"`
}

// 反射：在程序运行期对程序进行访问和修改的能力
func main() {
	str := `{"name": "周林", "age": 9000}`
	var p person
	json.Unmarshal([]byte(str), &p)
	fmt.Printf(p.Name, p.Age)

	var c = Cat{}
	reflectType(c)

	d := 100
	reflectSetValue1(&d)

}

type Cat struct {
}

// TypeOf 获取类型
func reflectType(x interface{}) {
	v := reflect.TypeOf(x)
	fmt.Printf("type:%v\n", v)
}

// ValueOf 原始的值信息
func reflectValue(x interface{}) {
	v := reflect.ValueOf(x)
	fmt.Printf("%v\n", v.Int())
}

// 反射中类型分为两种：类型(Type)和种类（Kind）
// 因为可以使用type关键字构造很多自定义类型，Kind指底层数据类型，

// Type类型为Cat，Kind类型为struct
//type Cat struct {
//
//}

func Test(x interface{}) {
	v := reflect.TypeOf(x)
	k := v.Kind()

	switch k {
	case reflect.Int64:
		fmt.Println("")
	case reflect.Int:
	}

}

// 通过反射设置变量的值
func reflectSetValue1(x interface{}) {
	// 反射种使用Elem()方法获取指针对应的值
	v := reflect.ValueOf(x)
	//if v.Kind() == reflect.Int64 {
	//	v.SetInt(200)
	//}
	if v.Elem().Kind() == reflect.Int64 {
		v.Elem().SetInt(200)
	}
}

//isNil() 报告v持有的值是否为nil，v持有的值分类必须是通道、函数、接口、映射、指针、切片之一（引用类型）
// 否则会panic
