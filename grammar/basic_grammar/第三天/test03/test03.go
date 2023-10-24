package main

import "fmt"

func main() {
	studentMap := make(map[string]map[string]int, 10)
	//创建内层map
	studentMap["Sam"] = make(map[string]int, 3)
	studentMap["Sam"]["id"] = 1
	studentMap["Sam"]["age"] = 28
	studentMap["Sam"]["num"] = 8665842

	studentMap["amy"] = make(map[string]int, 3)
	studentMap["amy"]["id"] = 2
	studentMap["amy"]["age"] = 33
	studentMap["amy"]["num"] = 111111
	for k, v := range studentMap {
		fmt.Println(k)
		if v["id"] == 1 {
			for k2, v2 := range v {
				fmt.Println(k2, v2)
			}
		}
	}

}
