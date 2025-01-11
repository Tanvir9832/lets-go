package main

import "fmt"

type Person struct {
	Name       string
	Age        int
	University string
	Department string
	Cgpa       float64
}

func main() {
	var tanvir Person
	fmt.Println(tanvir) // { 0  0}
	tanvir.Name = "Tanvir"
	tanvir.Age = 22
	tanvir.University = "Green University of Bangladesh"
	tanvir.Department = "CSE"
	tanvir.Cgpa = 1.50

	fmt.Println(tanvir) // {Tanvir 22 Green University of Bangladesh CSE 1.5}

	// Another way to declare struct
	Anik := Person{
		Name:       "Tanvir",
		Age:        22,
		University: "Green University of Bangladesh",
		Department: "CSE",
		Cgpa:       1.50,
	}

	fmt.Println(Anik) // {Tanvir 22 Green University of Bangladesh CSE 1.5}



	type address struct {
		
	}
}
