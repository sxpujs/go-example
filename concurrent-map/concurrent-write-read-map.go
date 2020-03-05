package main

import (
	"encoding/json"
	"fmt"
	"strconv"
)

func main() {

	m := map[string]string{
		"name": "andy",
		"age":  "30",
	}

	write := make(chan bool)
	go func() {
		for i := 0; i < 10000; i++ {
			m["idx"] = strconv.Itoa(i)
		}
		write <- true
	}()

	read := make(chan bool)
	go func() {
		for i := 0; i < 10000; i++ {
			json.Marshal(m)
		}
		read <- true
	}()

	<-write
	<-read

	fmt.Println("Finish")
}
