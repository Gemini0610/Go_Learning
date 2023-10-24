1、数组定义 数组要带长度和存储的类型
   var a [3] int 
   数组初始化
   var a[3]int{1,2,3}
   var b= [...]bool{true,false} 未规定长度自动填充
   var c = [5]{1:"哈哈哈",3:"ddd"} 支持索引赋值
   多维数组
   var a [3][2]int ={
    {1,2},
    {2,3},
    {33,4}
   }
   遍历 for index, value := range b{
    fmt.Pringln(index, value)
   }

2、切片：是语言层面上对数组的封装
   切片与数组的区别；切片没有固定的长度
   数组是值类型，切片是引用类型(初始值是申请一个内存空间，所以是nil)
   大小len是指当前切片中元素的个数
   容量cap是指底层数组能容下多少个元素
   make创建切片 make(切片类型，大小，容量)
   append函数追加切片元素 必须要用变量接收

3、map 引用类型
   map专门来存储键值对
   判断某个键是否存在 v, ok := m[key]

4、函数
   无参 func f1(){}
   多个参数 func f2(name1, name2, string){}
   可变参数 func f3(name ...string){}

   匿名函数 var f = func (name string){}
            通过变量调用 或者  命名并调用
   
   闭包 函数调用了外部的变量
   defer 延迟

5、类型别名只存在于代码编译过程中

6、结构体 定义type ___ struct{}
   实例化  (1)var 名称 = 结构体名字{}
          (2)var 名称 = new(结构体名字)   此时指向的是内存地址
   匿名结构体 字段没有名字 以类型进行定义，但不能有多个