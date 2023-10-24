package split

import "strings"

//定义一个切割字符串的包
//Split 用sep分割s
//a:b:c :---->[a b c]
func Split(s, sep string) (result []string) {
	ok := strings.Contains(s, sep)
	if !ok { //s 不包含sep，返回一个只有s的切片
		return []string{s}
	}
	index := strings.Index(s, sep)
	for index >= 0 {
		result = append(result, s[:index])
		s = s[index+len(sep):] //使用len(sep)获取sep的长度，这里是为了考虑多字符的情况
		index = strings.Index(s, sep)
	}
	result = append(result, s)
	return
}
