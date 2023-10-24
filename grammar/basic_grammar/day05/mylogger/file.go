package mylogger

import (
	"fmt"
	"os"
	"path"
	"time"
)

//往文件中写日志的代码

// FileLogger 结构体
type FileLogger struct {
	fileName string
	filePath string
	file     *os.File
	errFile  *os.File
}

// 文件日志结构体的构造函数
func NewFileLogger(fileName, filePath string) *FileLogger {
	f1 := &FileLogger{
		fileName: fileName,
		filePath: filePath,
	}
	f1.initFile() //根据文件路径与文件名，将文件句柄赋值给文件结构体对应的字段
	return f1
}

// 将打开的文件，赋值给结构体
func (f *FileLogger) initFile() {
	logName := path.Join(f.filePath, f.fileName)
	//打开文件
	fileObj, err := os.OpenFile(logName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0664)
	if err != nil {
		panic(fmt.Errorf("open document is failure, %v", logName, err))
	}
	f.file = fileObj
	//打开错误文件
	errlogName := fmt.Sprintf("%s.err", logName)
	//打开文件
	errFileObj, err := os.OpenFile(errlogName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0664)
	if err != nil {
		panic(fmt.Errorf("open document is failure, %v", errlogName, err))
	}
	f.errFile = errFileObj
}
func (f *FileLogger) Debug(format string, args ...interface{}) {
	//得到用户要记录的日志
	msg := fmt.Sprintf(format, args...)
	//日志格式[时间][文件:行号][函数名][日志级别] 日志信息
	timeStr := time.Now().Format("2006-01-01 15:04:05.000 ")
	fileName, line, funcName := getCallerInfo(3)
	logMsg := fmt.Sprintf("[%s][%s:%d][%s][%s]%s", timeStr, fileName, line, funcName, "debug", msg)
	//利用fmt将msg写入f.file文件中
	fmt.Fprintln(f.file, logMsg)
}
