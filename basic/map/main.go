package main

import "fmt"

func main() {
	map1 := make(map[string]string, 10)
	map1["a"] = "b"
	for idx, val := range map1 {
		fmt.Println(idx, val)
	}

	map1["c"] = "d"
	value, exists := map1["c"]
	if exists {
		fmt.Println(value)
	}

	myFuncMap := map[string]func() int{
		"func1": func() int { return 1 },
	}
	fmt.Println(myFuncMap)
}
