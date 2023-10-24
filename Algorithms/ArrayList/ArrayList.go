package Arraylist

import "errors"

//接口
type List interface {
	Size() int //数组大小
	//抓取第几个元素,返回接口类型，以及错误信息
	Get(index int) (interface{}, error)
	//设置修改数据，newval标注新数据的类型
	Set(index int, newval interface{}) error
	//插入数据
	Insert(index int, newval interface{}) error
	//追加数据
	Append(newval interface{}) error
	//清空
	Clear()
	//删除
	Delete(index int) error
	//返回字符串
	String() string
}

//数据结构，字符串，整数，实数
type Arraylist struct {
	//存储数据，需要是字符串就是字符串，属于泛型
	dataStore []interface{}
	theSize   int //数组的大小
}

//创建
func NewArrayList() *Arraylist {
	list := new(Arraylist)                      //初始化结构体
	list.dataStore = make([]interface{}, 0, 10) //开辟空间10个
	list.theSize = 0                            //初始化
	return list
}

func (list *Arraylist) Size() int {
	return list.theSize //返回数组大小
}

//抓取数据
func (list *Arraylist) Get(index int) (interface{}, error) {
	if index < 0 || index > list.Size() {
		//索引大于数组大小或小于0，报错
		return nil, errors.New("索引越界")
	}
	return list.dataStore[index], nil
}

//追加数据
func (list *Arraylist) Append(newval interface{}) error {

}
