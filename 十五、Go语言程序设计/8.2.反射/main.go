package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

/*
不应滥用反射：
1.反射中的类型错误会在真正运行的时候才引发panic,可能是在代码写完的很长时间之后
2.大量使用发射的代码通常难以理解
3.反射性能低下，基于反射实现的代码要比一般代码慢一到两个数量级
 */
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
	t := reflect.TypeOf(x)
	// Name()返回类型名，而Kind返回类型所属的种类
	// 例如设置自定义类型： type myInt int
	// 则Name为: myInt, Kind为int(底层数据类型)
	// 如果为切片、函数、数组、Map等类型变量，Name()返回空
	fmt.Printf("type:%v,%v\n", t.Name(), t.Kind())
}

// ValueOf 原始的值信息
func reflectValue(x interface{}) {
	v := reflect.ValueOf(x)
	// 将值以interface类型返回，返回的值可以通过类型断言转换为指定类型
	fmt.Println(v.Interface())
	fmt.Printf("%v\n", v.Int())
}

// 反射中类型分为两种：类型(Type)和种类（Kind）
// 因为可以使用type关键字构造很多自定义类型，Kind指底层数据类型，

// Type类型为Cat，Kind类型为struct
//type Cat struct {
//
//}

func Test(x interface{}) {
	t := reflect.TypeOf(x)
	k := t.Kind()

	switch k {
	case reflect.Int64:
		fmt.Println("")
	case reflect.Int:
	}

}

func reflectValue2(x interface{}) {
	v := reflect.ValueOf(x)
	k := v.Kind()
	switch k {
	case reflect.Int, reflect.Int8, reflect.Int64:
		fmt.Println(v.Int())
	case reflect.String:
		fmt.Println(v.String())
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

//func (v Value) IsNil() bool 报告v持有的值是否为nil，v持有的值分类必须是通道、函数、接口、映射、指针、切片之一（引用类型）; 否则会panic
//func (v Value) IsValid() bool 返回v是否持有一个值。如果v是Value零值会返回假，此时v除了IsValid、String、Kind之外的方法都会导致panic。

// IsNil用于判断指针是否为空， IsValid判断值是否有效
func NilAndValid(x interface{}) {
	v := reflect.ValueOf(x)
	fmt.Println(v.IsNil(), v.IsValid())
}

// 获取结构体信息
type student struct {
	Name  string `json:"name"`
	Score int    `json:"score"`
}

func GetStructField(x interface{}) {
	v := reflect.ValueOf(x)
	k := v.Kind()
	if k == reflect.Struct {
		t := reflect.TypeOf(x)
		for i := 0; i < t.NumField(); i++ {
			field := t.Field(i)
			fmt.Println(field.Name, field.Index, field.Type, field.Tag.Get("json"))
		}
		if s, ok := t.FieldByName("Score"); ok {
			fmt.Println(s.Type, s.Name, s.Tag.Get("json"))
		}
	}
}

func PrintMethod(x interface{}) {
	t := reflect.TypeOf(x)
	v := reflect.ValueOf(x)
	fmt.Println(t.NumMethod())
	for i := 0; i < v.NumMethod(); i++ {
		methodType := v.Method(i).Type()
		fmt.Printf("method:%s\n", t.Method(i).Name)
		fmt.Printf("method:%s\n", methodType)
		// 通过反射调用方法传递的参数必须是 []reflect.Value 类型
		args := make([]reflect.Value, 0)
		v.Method(i).Call(args)
	}
}
//isNil() 报告v持有的值是否为nil，v持有的值分类必须是通道、函数、接口、映射、指针、切片之一（引用类型）
// 否则会panic
