1.1
go build hello.go 会生成一个可执行的二进制文件 .exe,可以随时运行

go语言的包，类似于其他语言中的库或者模块，一个包由位于单个目录下的一个或多个.go源代码组成
每个源文件，都以一条package声明语句开始，表示该文件属于哪个包，紧接着import

main包定义了一个独立可执行的程序，不是一个库，在mian包里的main函数特殊:它是整个程序执行的入口

import声明必须跟在文件的package声明后，声明语句 func函数、var变量、const常量、type类型
