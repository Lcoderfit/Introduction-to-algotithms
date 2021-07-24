package mylogger

// 往文件中写日志
type FileLogger struct {
	Level       LogLevel
	filePath    string // 日志文件保存路径
	fileName    string // 日志文件保存文件名
	maxFileSize int64  // 最大文件大小，超过就切割文件
}

func NewFileLogger(levelStr, fp, fn string, maxSize int64) *FileLogger {
	logLevel, err := parseLogLevel(levelStr)
	if err != nil {
		panic(err)
	}
	return &FileLogger{
		Level:       logLevel,
		filePath:    fp,
		fileName:    fn,
		maxFileSize: maxSize,
	}

}

func main() {

}
