package main

import (
	"fmt"

	_ "github.com/miles0wu/go101/basic/init/a"
	_ "github.com/miles0wu/go101/basic/init/b"
)

func init() {
	fmt.Println("main init")
}

// output:
// b init
// a init
// main init
// main
func main() {
	fmt.Println("main")
}
