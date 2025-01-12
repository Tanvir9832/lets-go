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

	type Address struct {
		Street  string
		City    string
		State   string
		Zip     int
		country string
	}

	type Contract struct {
		email   string
		phone   string
		address Address
	}

	type Employee struct {
		name     string
		jobTitle string
		contract Contract
	}

	firstEmployee := Employee{
		name:     "Tanvir",
		jobTitle: "Software Engineer",
		contract: Contract{
			email: "Tanvir@gmail.com",
			phone: "+8801739783323",
			address: Address{
				Street:  "Dhanmondi",
				City:    "Dhaka",
				State:   "Dhaka",
				Zip:     1205,
				country: "Bangladesh",
			},
		},
	}
	firstEmployee.contract.address.City = "Gaza"
	firstEmployee.contract.address.country = "Palestine"
	fmt.Println(firstEmployee)
}
