package main

import "fmt"

func main() {
	var arr [5]int
	arr[0] = 10
	arr[1] = 20
	arr[2] = 30
	arr[3] = 40
	arr[4] = 50

	var name [5]string
	name[0] = "Md."
	name[1] = "Tanvir"
	name[2] = "Alam"
	name[3] = "Anik"

	fmt.Println("Print the whole array name : ", name)

	// var arr2 = [5]int{10, 20, 30, 40, 50}

	// var arr3 = [...]int{10, 20, 30, 40, 50}

	// var arr4 = [5]string{3: "Chris", 4: "Ron"}

	// var arr5 = [5]int{1: 10, 3: 30}

	var arrs []int

	for i := 0; i < 20; i++ {
		arrs = append(arrs, i)
	}

	fmt.Println("Length of arrs : ", len(arrs))   //output: 20 Explanation: The length of the slice is 20 means it has 20 elements
	fmt.Println("Capacity of arrs : ", cap(arrs)) //output: 32 Explanation: The capacity of the slice is 32 means it can hold 32 elements

	//!Number of array

	var arrWithSize = make([]int, 5)
	fmt.Println("Total array is : ", arrWithSize) // output: [0 0 0 0 0]
	arrWithSize = append(arrWithSize, 10)
	arrWithSize[3] = 5
	fmt.Println("Total array is : ", arrWithSize) // output: [0 0 0 5 0 10]

	var arrWithSize2 []int
	fmt.Println("Total array is : ", arrWithSize2) // output: []
	var arrWithSize3 [10]int
	fmt.Println("Total array is : ", arrWithSize3) // output: [0 0 0 0 0 0 0 0 0 0]

	//! String of array
	var names [10]string
	fmt.Println("Total array is : ", names) // output: [         ]
	names[5] = "Md."
	fmt.Println("Total array is : ", names)    // output: [         Md.      ]
	fmt.Printf("Total array is : %q\n", names) // output: ["         " "Md.      " "         " "         " "         " "         " "         " "         " "         " "         "]

	//! Boolean of array
	var isExist [5]bool
	fmt.Println("Total array is : ", isExist) // output: [false false false false false]
	isExist[3] = true
	fmt.Println("Total array is : ", isExist) // output: [false false false true false]

	//! Float of array
	var floatArr [5]float64
	fmt.Println("Total array is : ", floatArr) // output: [0 0 0 0 0]
	floatArr[3] = 5.5
	fmt.Println("Total array is : ", floatArr) // output: [0 0 0 5.5 0]

	//! Array of array
	var arrOfArr [3][3]int
	fmt.Println("Total array is : ", arrOfArr) // output: [[0 0 0] [0 0 0] [0 0 0]]
	arrOfArr[1][1] = 5
	fmt.Println("Total array is : ", arrOfArr) // output: [[0 0 0] [0 5 0] [0 0 0]]

	//! Array of slice
	var arrOfSlice [3][]int
	fmt.Println("Total array is : ", arrOfSlice) // output: [[0 0 0] [0 0 0] [0 0 0]]
	arrOfSlice[1] = append(arrOfSlice[1], 5)
	fmt.Println("Total array is : ", arrOfSlice) // output: [[0 0 0] [5] [0 0 0]]

	//! Slice of slice
	var sliceOfSlice [][]int
	fmt.Println("Total array is : ", sliceOfSlice) // output: []
	sliceOfSlice = append(sliceOfSlice, []int{1, 2, 3})
	fmt.Println("Total array is : ", sliceOfSlice) // output: [[1 2 3]]

	//! Slice of array
	var sliceOfArray [][3]int
	fmt.Println("Total array is : ", sliceOfArray) // output: []
	sliceOfArray = append(sliceOfArray, [3]int{1, 2, 3})
	fmt.Println("Total array is : ", sliceOfArray) // output: [[1 2 3]]

	//! Array of map
	var arrOfMap [3]map[string]int
	fmt.Println("Total array is : ", arrOfMap) // output: [map[] map[] map[]]
	arrOfMap[1] = make(map[string]int)
	arrOfMap[1]["key"] = 5
	fmt.Println("Total array is : ", arrOfMap) // output: [map[] map[key:5] map[]]

	//! Map of array
	var mapOfArray map[string][3]int
	fmt.Println("Total array is : ", mapOfArray) // output: map[]
	mapOfArray = make(map[string][3]int)
	mapOfArray["key"] = [3]int{1, 2, 3}
	fmt.Println("Total array is : ", mapOfArray) // output: map[key:[1 2 3]]

	//! Map of slice
	var mapOfSlice map[string][]int
	fmt.Println("Total array is : ", mapOfSlice) // output: map[]
	mapOfSlice = make(map[string][]int)
	mapOfSlice["key"] = append(mapOfSlice["key"], 1)
	fmt.Println("Total array is : ", mapOfSlice) // output: map[key:[1]]

	//! Slice of map
	var sliceOfMap []map[string]int
	fmt.Println("Total array is : ", sliceOfMap) // output: []
	sliceOfMap = append(sliceOfMap, make(map[string]int))
	sliceOfMap[0]["key"] = 5
	fmt.Println("Total array is : ", sliceOfMap) // output: [map[key:5]]

	//! Slice of struct
	type person struct {
		name string
		age  int
	}
	var sliceOfStruct []person
	fmt.Println("Total array is : ", sliceOfStruct) // output: []
	sliceOfStruct = append(sliceOfStruct, person{name: "Tanvir", age: 25})
	fmt.Println("Total array is : ", sliceOfStruct) // output: [{Tanvir 25}]

	//! Array of struct
	var arrOfStruct [3]person
	fmt.Println("Total array is : ", arrOfStruct) // output: [{ 0} { 0} { 0}]
	arrOfStruct[1] = person{name: "Tanvir", age: 25}
	fmt.Println("Total array is : ", arrOfStruct) // output: [{ 0} {Tanvir 25} { 0}]

	//! Map of struct
	var mapOfStruct map[string]person
	fmt.Println("Total array is : ", mapOfStruct) // output: map[]
	mapOfStruct = make(map[string]person)
	mapOfStruct["key"] = person{name: "Tanvir", age: 25}
	fmt.Println("Total array is : ", mapOfStruct) // output: map[key:{Tanvir 25}

	//! Struct of struct
	type address struct {
		city string
		name string
	}
	type personWithAddress struct {
		name    string
		age     int
		address address
	}
	var structOfStruct personWithAddress
	fmt.Println("Total array is : ", structOfStruct) // output: { 0 { }}
	structOfStruct = personWithAddress{name: "Tanvir", age: 25, address: address{city: "Dhaka", name: "Mirpur"}}
	fmt.Println("Total array is : ", structOfStruct) // output: {Tanvir 25 {Dhaka Mirpur}}

	//! Struct of slice
	type personWithAddressSlice struct {
		name    string
		age     int
		address []address
	}
	var structOfSlice personWithAddressSlice
	fmt.Println("Total array is : ", structOfSlice) // output: { 0 []}
	structOfSlice = personWithAddressSlice{name: "Tanvir", age: 25, address: []address{{city: "Dhaka", name: "Mirpur"}}}
	fmt.Println("Total array is : ", structOfSlice) // output: {Tanvir 25 [{Dhaka Mirpur}]}

	//! Struct of map
	type personWithAddressMap struct {
		name    string
		age     int
		address map[string]address
	}
	var structOfMap personWithAddressMap
	fmt.Println("Total array is : ", structOfMap) // output: { 0 map[]}
	structOfMap = personWithAddressMap{name: "Tanvir", age: 25, address: map[string]address{"home": {city: "Dhaka", name: "Mirpur"}}}
	fmt.Println("Total array is : ", structOfMap) // output: {Tanvir 25 map[home:{Dhaka Mirpur}]}

	//! array methods
	var arrs2 = []int{1, 2, 3, 4, 5}
	fmt.Println("Total array is : ", arrs2) // output: [1 2 3 4 5]
	arrs2 = append(arrs2, 6)
	fmt.Println("Total array is : ", arrs2) // output: [1 2 3 4 5 6]
	arrs2 = append(arrs2, 7, 8, 9)
	fmt.Println("Total array is : ", arrs2) // output: [1 2 3 4 5 6 7 8 9]
	arrs2 = append(arrs2, []int{10, 11, 12}...)
	fmt.Println("Total array is : ", arrs2) // output: [1 2 3 4 5 6 7 8 9 10 11 12]
	arrs2 = append(arrs2[:3], arrs2[4:]...)
	fmt.Println("Total array is : ", arrs2)      // output: [1 2 3 5 6 7 8 9 10 11 12]
	fmt.Println("Total array is : ", arrs2[2:5]) // output: [3 5 6]
	fmt.Println("Total array is : ", arrs2[:5])  // output: [1 2 3 5 6]
	fmt.Println("Total array is : ", arrs2[5:])  // output: [7 8 9 10 11 12]
	fmt.Println("Total array is : ", arrs2[:])   // output: [1 2 3 5 6 7 8 9 10 11 12]
	fmt.Println("Total array is : ", arrs2[5:5]) // output: []

}
