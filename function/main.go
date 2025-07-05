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

//declaring a struct takes name, age
type Person struct {
	Name string
	Age  int
}

//greet function takes a person and returns a string
func (p Person) greet() {
	fmt.Println("Hello, my name is", p.Name)
}

func (p *Person) getName() string {
	return p.Name
}

func (p *Person) setName(name string) {
	p.Name = name
}

func (p *Person) getAge() int {
	return p.Age
}

func (p *Person) setAge(age int) {
	p.Age = age
}

func (p Person) getPerson() Person {
	return p
}

func init() {
	fmt.Println("Am I the first one to print even before the main function ?")
}

func main() {
	fmt.Println(allPossible(10, 5))
	person := Person{
		Name: "Tanvir",
		Age:  22,
	}

	person.greet()
	fmt.Println(person.getName())
	person.setName("Md. Tanvir Alam Anik")
	fmt.Println(person.getName())
	fmt.Println(person.getAge())
	person.setAge(9)
	fmt.Println(person.getAge())
	fmt.Println(person.getPerson())

	fmt.Println(person)
	fmt.Println(person.Name)
	person.greet()
}
