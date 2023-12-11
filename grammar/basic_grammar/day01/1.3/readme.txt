1.3 
1、map存储了键值对,键任意类型，只要其值能用==运算符比较
                  值也可以是任意类型
    counts[input.Text()]++等价于
    line := input.Text()
    counts[line] = counts[line] + 1

2、bufio包处理输入和输出,方便高效
   input.Scan()读取下一行 input.Text()获取读取的内容
   %d 表示以十进制形式打印一个整型操作数
   %s 表示把字符串型操作数的值展开

3、%d          十进制整数
   %x, %o, %b  十六进制，八进制，二进制整数。
   %f, %g, %e  浮点数： 3.141593 3.141592653589793 3.141593e+00
   %t          布尔：true或false
   %c          字符（rune） (Unicode码点)
   %s          字符串
   %q          带双引号的字符串"abc"或带单引号的字符'c'
   %v          变量的自然形式（natural format）
   %T          变量的类型
   %%          字面上的百分号标志（无操作数）

4、\t 制表符 \n 换行符