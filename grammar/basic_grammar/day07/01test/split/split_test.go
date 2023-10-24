package split

import (
	"reflect"
	"testing"
)

// Test函数要以Test开头
func TestSplit(t *testing.T) {
	//获得的结果
	got := Split("a:b:c", ":")
	//期待的结果
	want := []string{"a", "b", "c"}
	//判断是否一致
	if ok := reflect.DeepEqual(got, want); !ok {
		t.Fatalf("期望得到:%v, 实际得到:%v", want, got)
	}
}
func TestMoreSplit(t *testing.T) {
	got := Split("abcd", "bc")
	want := []string{"a", "d"}
	if !reflect.DeepEqual(want, got) {
		t.Errorf("期望值:%v,真实值:%v", want, got)
	}
}

// 测试中文切片,使用了测试组
func TestChineseSplit(t *testing.T) {
	//定义一个测试用例类型
	type test struct {
		input string
		sep   string
		want  []string
	}
	//定义一个存储测试用例的切片
	tests := []test{
		{input: "a:b:c", sep: ":", want: []string{"a", "b", "c"}},
		{input: "a:b:c", sep: ",", want: []string{"a:b:c"}},
		{input: "沙和尚沙丘沙堆", sep: "沙", want: []string{"和尚", "丘", "堆"}},
	}
	//遍历切片，逐一执行测试用例
	for _, tc := range tests {
		got := Split(tc.input, tc.sep)
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("期望值:%v,真实值:%v", tc.want, got)
		}

	}
}

func TestSplit1(t *testing.T) {
	type test struct { // 定义test结构体
		input string
		sep   string
		want  []string
	}
	tests := map[string]test{ // 测试用例使用map存储
		"simple":      {input: "a:b:c", sep: ":", want: []string{"a", "b", "c"}},
		"wrong sep":   {input: "a:b:c", sep: ",", want: []string{"a:b:c"}},
		"more sep":    {input: "abcd", sep: "bc", want: []string{"a", "d"}},
		"leading sep": {input: "沙河有沙又有河", sep: "沙", want: []string{"河有", "又有河"}},
	}
	for name, tc := range tests {
		got := Split(tc.input, tc.sep)
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("name:%s expected:%#v, got:%#v", name, tc.want, got) // 将测试用例的name格式化输出
		}
	}
}
