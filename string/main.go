package main

import (
	"fmt"
	"strings"
)

func main() {
	//!string split
	elements := "Banana , Orange , Mango"
	fruits := strings.Split(elements, " , ")
	fmt.Println(fruits)

	//! normal assignment
	//! In assignment both all the variable point the same underlying data . But underlying data can not be changed, if we want to change the data then it will create a new string will new memory allocation

	normalClone := elements

	fmt.Println("Normal clone data --> " + normalClone)

	//! string clone
	//! Create a complete new string
	clone := strings.Clone(elements)
	fmt.Println("Clone method -->" + clone)

	//! string compare
	fmt.Println(strings.Compare("a", "b"))           //! output : -1	"small, large" -1
	fmt.Println(strings.Compare("a", "a"))           //! output : 0		"equal" 0
	fmt.Println(strings.Compare("b", "a"))           //! output : 1		"large, small" 1
	fmt.Println(strings.Compare("A", "a"))           //! output : -1  	"small, large" -1
	fmt.Println(strings.Compare("amader", "acader")) //! output : 1

	//!string contains
	fmt.Println(strings.Contains("sigma", "gma"))   //!output : true
	fmt.Println(strings.Contains("sigma", "bigma")) //!output : false
	fmt.Println(strings.Contains("", ""))           //!output : true

	//!string containsAny
	fmt.Println(strings.ContainsAny("sigma", "ihnb")) //!true  atleat one word matches (i)
	fmt.Println(strings.ContainsAny("sigma", "gnb"))  //! false  no elements matches

	//!string count
	s := "Md. Tanvir Alam Anik"
	count := strings.Count(s, "a")
	fmt.Println(count) //! output : 2
	count = strings.Count(s, "")
	fmt.Println(count) //!output : 21

	before, after, exist := strings.Cut(s, "A")
	fmt.Printf("before is = (%s) and after = (%s) exist = (%t) \n", before, after, exist) //! before is = (Md. Tanvir ) and after = (lam Anik) exist = (true)
	before, after, exist = strings.Cut(s, "")
	fmt.Printf("before is = (%s) and after = (%s) exist = (%t) \n", before, after, exist) //! before is = () and after = (Md. Tanvir Alam Anik) exist = (true)
	before, after, exist = strings.Cut(s, " ")
	fmt.Printf("before is = (%s) and after = (%s) exist = (%t) \n", before, after, exist) //! before is = (Md.) and after = (Tanvir Alam Anik) exist = (true)
	//!! CutSuffix (before) , CutPrefix  (after)

	//! TrimSpace
	s = "            Tanvir             Alam             "
	fmt.Println(strings.TrimSpace(s)) //!Tanvir             Alam

	//! Upper and Lower
	fmt.Println(strings.ToUpper(s)) //!TANVIR             ALAM
	fmt.Println(s)                  //!Tanvir             Alam
	fmt.Println(strings.ToLower(s)) //!tanvir             alam

	//! Fields
	fmt.Printf("Fields are: %q", strings.Fields("  tanvir    alam    anik  ")) //!Fields are: ["foo" "bar" "baz"]

	//! Index
	fmt.Println(strings.Index("chicken", "ken")) //! 4
	fmt.Println(strings.Index("chicken", "dmr")) //! -1

	//!Join
	fruits = []string{"mango", "banana", "orange"}
	fmt.Println(strings.Join(fruits, ", ")) //! mango, banana, orange

	//!Replace
	fmt.Println(strings.Replace("oink oink oink", "k", "ky", 2))      //!oinky oinky oink
	fmt.Println(strings.Replace("oink oink oink", "oink", "moo", -1)) //! moo moo moo

	//!ReplaceAll
	fmt.Println(strings.ReplaceAll("oink oink oink", "oink", "moo")) //! moo moo moo
}
