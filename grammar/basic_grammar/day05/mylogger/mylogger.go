package mylogger

//我的日志库
//level  自定义一个类型 代表日志级别
type Level uint16

//定义具体的日志级别常量
const (
	DebugLevel Level = iota
	InfoLevel
	WarningLevel
	ErrorLevel
	FetaLevel
)
