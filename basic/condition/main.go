package main

import (
	"fmt"
	"time"
)

func main() {
	a := 2

	if a > 0 {
		fmt.Println("a is greater than 0")
	}

	switch hour := time.Now().Hour(); {
	case hour < 12:
		fmt.Println("Good morning")
	case hour < 17:
		fmt.Println("Good afternoon")
	default:
		fmt.Println("Good evening")
	}
}
