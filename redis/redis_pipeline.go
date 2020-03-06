package main

import (
	"github.com/gomodule/redigo/redis"
	"log"
	"time"
)

var (
	c     redis.Conn
	err   error
	reply interface{}
)

func init() {
	c, err = redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	defer c.Close()
	withoutPipelining() // 256.53 ms
	withPipelining()    // 8.69 ms
}

func withoutPipelining() {
	defer timeTrack(time.Now(), "withoutPipelining")
	for i := 0; i < 10000; i++ {
		c.Do("PING")
	}
}

func withPipelining() {
	defer timeTrack(time.Now(), "withPipelining")
	c.Send("MULTI")
	for i := 0; i < 10000; i++ {
		c.Send("PING")
	}
	c.Do("EXEC")
}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}
