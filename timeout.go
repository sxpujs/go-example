package main

import (
	"fmt"
	"time"
)

func main() {

	ch := make(chan bool)

	go func() {
		time.Sleep(time.Second * 2)
		fmt.Println("hello")
		ch <- true
	}()

	select {
	case res := <-ch:
		fmt.Println(res)
	case <-time.After(time.Second * 1):
		fmt.Println("timeout")
	}

	fmt.Println("exit")
}
