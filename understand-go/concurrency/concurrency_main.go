package main

import (
	"fmt"
	"sync"
	"time"
)

var counter = 0
var lock sync.Mutex

func main() {
	for i := 0; i < 2; i++ {
		go incr()
	}
	time.Sleep(time.Millisecond * 10)
	lock.Lock()
}

func incr() {
	lock.Lock()
	counter++
	lock.Unlock()
	fmt.Println(counter)
}
