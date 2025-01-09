package main

import "fmt"

func main() {
	slice := []int{}
	fmt.Printf("Total slice is : %T\n", slice) // output: []int
	fmt.Println("Total slice is : ", slice)    // output: []

	sliceWithMake := make([]int, 5, 10)
	fmt.Printf("Total slice is : %v", sliceWithMake) // output: [0 0 0 0 0]

	sliceWithInitialValue := []int{1, 2, 3, 4, 5}
	fmt.Printf("Total slice is : %v", sliceWithInitialValue) // output: [1 2 3 4 5]

	//! Most of the thing coverd in array/main.go

}
