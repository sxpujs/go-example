package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	sm := SafeMap{m: map[int]int{}}

	for i := 0; i < 10000; i++ {
		go writeMap(&sm, i, i)
		go readMap(&sm, i)
	}
	time.Sleep(time.Second * 5)
	fmt.Println("quit")
}

type SafeMap struct {
	m map[int]int
	mux sync.RWMutex
}

func readMap(sm *SafeMap, key int) int {
	sm.mux.RLock()
	defer sm.mux.RUnlock()
	//fmt.Println("read map, key=", sm.m[key])
	return sm.m[key]
}

func writeMap(sm *SafeMap, key int, value int) {
	sm.mux.Lock()
	//fmt.Println("write map, key=", value)
	sm.m[key] = value
	sm.mux.Unlock()
}