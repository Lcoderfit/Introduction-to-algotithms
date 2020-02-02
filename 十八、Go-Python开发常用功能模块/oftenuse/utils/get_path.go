package utils

import (
	"runtime"
)

// n == 0, 返回runtime.Caller所在的文件（Get_path.go）的路径
// n == 1, 返回调用GetPath函数的文件（main.go）的路径
func GetPath(n int) string {
	// 返回命令（go run main.go）执行的路径
	//path, err := os.Getwd()
	//if err != nil {
	//	log.Fatal("GetPath Error: ", err)
	//	return ""
	//}
	if n != 1 && n != 0 {
		return ""
	}
	// 参数为0，filePath为runtime.Caller()函数所在的文件路径
	// 参数为1，filePath为调用runtime.Caller()所在的函数（GetPath）的文件路径
	_, filePath, _, ok := runtime.Caller(n)
	if !ok {
		return ""
	}

	return filePath
}
