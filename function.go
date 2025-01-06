package main

import "fmt"

func SimpleFunction() {
	fmt.Println("Hello, World!")
	fmt.Println(allPossible(10, 5))
	fmt.Println(myFunction(10, "Hello"))
}

//Function to add two numbers
func Add(a, b int) int {
	return a + b
}

//Function to do all possible operations
func allPossible(num1, num2 int) (sum, mul, sub, div int) {
	sum = num1 + num2
	mul = num1 * num2
	sub = num1 - num2
	div = num1 / num2
	return
}

func myFunction(x int, y string) (result int, txt1 string) {
	result = x + x
	txt1 = y + " World!"
	return
}

func singleReturnValue(firstNumber, secondNumber int) int {
	return firstNumber + secondNumber
}

func singleReturnValue2(firstNumber, secondNumber int) (result int) {
	return firstNumber + secondNumber
}

func singleReturnValue3(firstNumber, secondNumber int) (result int) {
	result = firstNumber + secondNumber
	return
}
