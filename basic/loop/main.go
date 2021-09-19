package main

import "fmt"

func main() {
	for i := 1; i < 3; i++ {
		fmt.Println(i)
	}

	// while
	sum := 20
	for sum < 300 {
		sum += sum
		fmt.Println(sum)
	}

	// infinite loop
	i := 2
	for {
		i += i
		if i > 50 {
			break
		}
	}

	var myStrings = []string{"do", "you", "wanna", "build", "a", "snowman"}

	for idx, val := range myStrings {
		fmt.Print(val + " ")
		if idx == len(myStrings)-1 {
			fmt.Print("\n")
		}
	}
}
