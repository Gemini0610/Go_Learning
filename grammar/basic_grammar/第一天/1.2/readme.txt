1、os包以跨平台的方式，提供了一些与操作系统交互的函数和变量，
   程序的命令行参数可从os包的Args变量获取，os包外部使用os.Args访问该变量
   s[i]访问单个元素，s[m:n]获取子序列，len(s)序列元素数目，左闭右开，包含n-m个元素

2、i++ 与 i+=1 等价 j=i++非法， ++和 -- 只能放在变量名后面。

3、for initialization; condition;post {

}
    initialization 语句是可选的，在循环开始前执行，
    condition语句是布尔表达式，其值在每次循环迭代开始时计算，ture执行循环体
    post语句在循环体执行结束后执行，之后再次对condition求值，直到condition为false

4、 _ 空标识符，可用于任何语法需要变量名，但程序逻辑不需要的时候

5、 s:="" 短变量声明，只能用于函数内部，不能用于包变量
    var s stirng 依赖于字符串的默认初始化零值机制
    var s = ""用的少除非同时声明多个变量
    var s stirng = "" 显式声明，类型不同可以

6、连接涉及的数据量极大，可以使用strings包中的join函数