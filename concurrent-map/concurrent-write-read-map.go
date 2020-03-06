package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

func main() {

	param := map[string]string{
		"name": "andy",
		"age":  "30",
	}

	f := func(i int) {
		tempMap := param
		//tempMap := DeepCopyMap(m) // 使用深度拷贝可以解决并发读写map的问题
		tempMap["idx"] = strconv.Itoa(i)
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(10)))
		json.Marshal(tempMap)
	}

	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(idx int) {
			defer wg.Done()
			f(idx)
		}(i)
	}
	wg.Wait()

	fmt.Println("Finish")
}

// DeepCopyMap map[string]string 类型实现深拷贝
func DeepCopyMap(params map[string]string) map[string]string {
	result := map[string]string{}
	for k, v := range params {
		result[k] = v
	}
	return result
}
