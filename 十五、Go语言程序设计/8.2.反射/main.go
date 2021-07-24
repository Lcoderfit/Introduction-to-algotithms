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
