package main

import "fmt"

func main() {

	//! defer
	fmt.Println("I am the first one to print !")
	for i := 0; i < 5; i++ {
		fmt.Println(i)
		defer fmt.Println(i)
	}
	fmt.Println("I am the second one to print !")

	defer fmt.Println("Hello")
	fmt.Println("I am the third one to print !")

	defer fmt.Println("World")
	fmt.Println("I am the fourth one to print !")

	//! OUTPUT
	// I am the first one to print !
	// 0
	// 1
	// 2
	// 3
	// 4
	// I am the second one to print !
	// I am the first one to print !
	// I am the first one to print !
	// World
	// Hello
	// 4
	// 3
	// 2
	// 1
	// 0

}
