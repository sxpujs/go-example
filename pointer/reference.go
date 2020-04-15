package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	data := map[string]interface{}{}
	changeWorks1(data)
	fmt.Println(data) // map[a:1]
	changeWorks2(&data)
	fmt.Println(data) // map[a:1 b:2]
	changeWorks3(data)
	fmt.Println(data) // map[a:1 b:2 c:3]
	changeNotWorks1(data)
	fmt.Println(data) // map[a:1 b:2 c:3]
	changeNotWorks2(&data)
	fmt.Println(data) // map[a:1 b:2 c:3]
}

// 函数生效，传入的是值
func changeWorks1(dest interface{}) {
	result := map[string]string{
		"a": "1",
	}
	m := dest.(map[string]interface{})
	for k, v := range result {
		m[k] = v
	}
}

// 函数生效，传入的是引用
func changeWorks2(dest interface{}) {
	result := map[string]string{
		"b": "2",
	}
	m := dest.(*map[string]interface{})
	for k, v := range result {
		(*m)[k] = v
	}
}

func changeWorks3(dest map[string]interface{}) {
	dest["c"] = "3"
}

// 函数修改dest未生效，传入的是值
func changeNotWorks1(dest interface{}) {
	result := map[string]string{
		"d": "4",
	}
	dest = result
}

// 函数修改dest未生效，传入的是引用
func changeNotWorks2(dest interface{}) {
	result := map[string]string{
		"e": "5",
	}
	dest = &result
}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}
