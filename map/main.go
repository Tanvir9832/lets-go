package main

import "fmt"

func main() {
	// name : Grade

	var Grade = make(map[string]string)
	Grade["Tanvir"] = "A+"
	Grade["Alam"] = "A"
	Grade["Anik"] = "A-"
	Grade["anik"] = "B+"

	fmt.Println(Grade) // output: map[Alam:A Anik:A- Tanvir:A+ anik:A-]

	fmt.Println("Grade of Tanvir is : ", Grade["Tanvir"]) // output: Grade of Tanvir is :  A+
	fmt.Println("Grade of Nokey is : ", Grade["NoKey"])   // output: Grade of Nokey is :  ""

	delete(Grade, "Tanvir")
	fmt.Println(Grade) // output: map[Alam:A Anik:A- anik:A-]

	grade, exist := Grade["Tanvir"]
	fmt.Println("Grade of Tanvir is : ", grade, " and exist : ", exist) // output: Grade of Tanvir is :   and exist :  false

	grade, exist = Grade["Alam"]
	fmt.Println("Grade of Alam is : ", grade, " and exist : ", exist) // output: Grade of Alam is :  A  and exist :  true

	for key, value := range Grade {
		fmt.Println("Grade of ", key, " is : ", value)
	}

	person := map[string]map[string]string{
		"Tanvir": {"Grade": "A+", "Roll": "1"},
		"Alam":   {"Grade": "A", "Roll": "2"},
		"Anik":   {"Grade": "A-", "Roll": "3"},
		"anik":   {"Grade": "B+", "Roll": "4"},
	}
	fmt.Println(person)                    // output: map[Alam:map[Grade:A Roll:2] Anik:map[Grade:A- Roll:3] Tanvir:map[Grade:A+ Roll:1] anik:map[Grade:B+ Roll:4]]
	fmt.Println(person["Tanvir"]["Grade"]) // output: A+
	fmt.Println(person["Alam"]["Roll"])    // output: 2

	person2 := map[string]int{
		"Tanvir": 1,
		"Alam":   2,
		"Anik":   3,
		"anik":   4,
	}
	fmt.Println(person2) // output: map[Alam:2 Anik:3 Tanvir:1 anik:4]

}
