package main

import (
	"GoLearn/mylogger"
)

func main() {
	//fileObj, err := Open("test.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	//if err != nil {
	//	fmt.Printf("open file failed, err: %s\n", err)
	//	return
	//}
	//// 输出到文件
	//// 重定向到标准输出：log.SetOutput(os.Stdout)
	//log.SetOutput(fileObj)
	//for {
	//	log.Println("这是一条测试日志")
	//	time.Sleep(3 * time.Second)
	//}

	log := mylogger.NewLog("error")
	log.Debug("debug级别日志")
	log.Trace("Trace级别日志")
	log.Info("Info级别日志")
	log.Warning("Warning级别日志")
	log.Error("Error级别日志")
	log.Panic("Panic级别日志")
}
