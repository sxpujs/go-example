package main

import (
	"fmt"
	"github.com/chasex/redis-go-cluster"
	"os"
	"time"
)

func main() {
	cluster, err := redis.NewCluster(
		&redis.Options{
			StartNodes:   []string{"127.0.0.1:30001", "127.0.0.1:30002", "127.0.0.1:30003"},
			ConnTimeout:  50 * time.Millisecond,
			ReadTimeout:  50 * time.Millisecond,
			WriteTimeout: 50 * time.Millisecond,
			KeepAlive:    16,
			AliveTime:    60 * time.Second,
		})
	if err != nil {
		fmt.Println("err=", err)
		os.Exit(1)
	}
	cluster.Do("SET", "foo", "bar")
	cluster.Do("INCR", "mycount", 1)
	cluster.Do("LPUSH", "mylist", "foo", "bar")
	cluster.Do("HMSET", "myhash", "f1", "foo", "f2", "bar")
}
