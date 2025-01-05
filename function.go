package main

import "fmt"

func SimpleFunction() {
	fmt.Println("Hello, World!")
}

//Function to add two numbers
func Add(a, b int) int {
	return a + b
}

func myFunction(x int, y string) (result int, txt1 string) {
	result = x + x
	txt1 = y + " World!"
	return
}
