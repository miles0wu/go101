package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// similar to stack
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
	// output: 3 2 1

	loop()
	time.Sleep(time.Second)
}

func loop() {
	lock := sync.Mutex{}
	for i := 0; i < 3; i++ {
		go func(i int) {
			lock.Lock()
			// do something...
			fmt.Println("loop", i)

			defer lock.Unlock()
		}(i)
	}
}
