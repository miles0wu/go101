package main

import "fmt"

func main() {
	myArray := [5]int{1, 2, 3, 4, 5}

	mySlice := myArray[1:4]
	fmt.Println(mySlice)

	mySlice = deleteItem(mySlice, 1)
	fmt.Println(mySlice)

	// new return pointer
	newSlice1 := new([]int)
	fmt.Println(newSlice1)

	// make return value
	newSlice2 := make([]int, 0)
	newSlice3 := make([]int, 10)
	newSlice4 := make([]int, 10, 20)
	fmt.Println(newSlice2)
	fmt.Println(newSlice3)
	fmt.Println(newSlice4)

	forslice := []int{1, 2, 3, 4, 5}
	for _, val := range forslice {
		val *= 2
	}
	fmt.Println(forslice)

	for idx := range forslice {
		forslice[idx] *= 2
	}

	fmt.Println(forslice)

}

func deleteItem(slice []int, index int) []int {
	return append(slice[:index], slice[index+1:]...)
}
