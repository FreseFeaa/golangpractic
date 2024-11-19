package main

import (
	"fmt"
	"sync"
	"time"
)

var schetchik int
var mutex sync.Mutex

func incr() {
	mutex.Lock()
	schetchik++
	mutex.Unlock()
}

func main() {
	for i := 0; i < 5; i++ {
		go incr()
	}

	time.Sleep(time.Second)

	fmt.Println("Schetchik itog:", schetchik)
}
