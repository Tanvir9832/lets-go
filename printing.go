package main

import "fmt"

func Printing() {
	// Different Printing in go
	fmt.Print("Hello, World!")

	// Println
	fmt.Println("Hello, World!")

	// Printf
	fmt.Printf("Hello, World!")

	// Println with formatting
	fmt.Println("Hello, World!", "This is a new line")

	// Printf with formatting
	fmt.Printf("Hello, World! %s", "This is a new line")

	// Println with formatting
	fmt.Println("Hello, World!", "This is a new line", "This is another new line")

	name, age := "Tanvir", 87

	fmt.Printf("My first name is %v and last name is %v \nMy name type is (%T) and my age type is (%T)", name, age, name, age)

	fraction := 3.14159
	fmt.Printf("Fraction is = (%f)", fraction)   //Fraction is = (3.141590)
	fmt.Printf("Fraction is = (%.2f)", fraction) //Fraction is = (3.14)
}
