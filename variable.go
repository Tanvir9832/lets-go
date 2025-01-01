package main

import (
	"fmt"
	"strconv"
)

func Variable() {
	var name string = "John Doe"
	fmt.Println(name)
	var age = 25
	fmt.Println(age)
	// age = "Tanvir"     Error: cannot use "Tanvir" (type string) as type int in assignment
	// age = 25.5         Error: cannot use 25.5 (type float64) as type int in assignment
	// age = true         Error: cannot use true (type bool) as type int in assignment

	var isTrue bool = false
	fmt.Println(isTrue)

	var height float64 = 5.8
	fmt.Println(height)

	const pi = 3.1416
	fmt.Println(pi)
	// pi = 3.14159 Error: cannot assign to pi

	//Multiple variable declaration
	var (
		roll  int    = 101
		dept  string = "CSE"
		batch        = 2015
	)

	fmt.Println(roll)
	fmt.Println(dept)
	fmt.Println(batch)

	// Short hand declaration
	rollNo := 102
	fmt.Println(rollNo)
	deptName := "EEE"
	fmt.Println(deptName)
	batchNo := 2016
	fmt.Println(batchNo)

	// Multiple variable declaration using short hand
	rollNo1, deptName1, batchNo1 := 103, "BBA", 2017
	fmt.Println(rollNo1)
	fmt.Println(deptName1)
	fmt.Println(batchNo1)

	// Zero value
	var a int
	var b string
	var c float64
	var d bool

	fmt.Println("A = ", a) //output: A = 0
	fmt.Println("B = ", b) //output: B =
	fmt.Println("C = ", c) //output: C = 0
	fmt.Println("D = ", d) //output: D = false

	// Zero value using short hand
	var e int
	f := ""
	g := 0.0
	h := false

	fmt.Println("E = ", e) //output: E = 0
	fmt.Println("F = ", f) //output: F =
	fmt.Println("G = ", g) //output: G = 0
	fmt.Println("H = ", h) //output: H = false

	// Zero value using short hand
	i, j, k, l := 0, "", 0.0, false

	fmt.Println("I = ", i) //output: I = 0
	fmt.Println("J = ", j) //output: J =
	fmt.Println("K = ", k) //output: K = 0
	fmt.Println("L = ", l) //output: L = false

	// Type conversion
	var m int = 10
	var n float64 = float64(m)
	fmt.Println(n)

	var o float64 = 20.5
	var p int = int(o)
	fmt.Printf("%T \n", o) //output: float64
	fmt.Printf("%T\n", p)  //output: int

	// var q string = "Hello"
	// var r int = int(q)     Error: cannot convert q (type string) to type int

	// var s string = "25"
	// var t int = int(s)      Error: cannot convert s (type string) to type int

	// var u string = "25"
	// var v int = int(u)      Error: cannot convert u (type string) to type int

	// var w string = "25"
	// var x int = int(w)      Error: cannot convert w (type string) to type int

	// Type conversion using short hand
	u := "25"
	v, _ := strconv.Atoi(u)
	fmt.Printf("%T\n", u) //output: string
	fmt.Printf("%T\n", v) //output: int

	// Type conversion using short hand
	w := 25
	x := strconv.Itoa(w)
	fmt.Printf("%T\n", w) //output: int
	fmt.Printf("%T\n", x) //output: string
}
