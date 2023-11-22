package main

import (
	"fmt"
)

// public variable can be created by first letter capital like :
var Token = "Rohit Varshney" //Public

func main() {
	var userName string = "Rohit"
	fmt.Println(userName)
	fmt.Printf("var is of type %T\n", userName)

	var boolCheck bool = true
	fmt.Println(boolCheck)
	fmt.Printf("var is of type %T\n", boolCheck)

	var smallVal uint = 125
	fmt.Println(smallVal)
	fmt.Printf("var is of type %T \n", smallVal)

	// Default value of int  = 0
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type %T \n", anotherVariable)

	// implicit type

	var website = "www.google.com"
	fmt.Println(website)
	fmt.Printf("Type %T \n", website)

	// no var type
	numberCount := 454345.22
	numberCount = 23
	fmt.Println(numberCount)
	fmt.Printf("Type is :%T \n", numberCount)

	fmt.Println(Token)

	// Multiple varsiables

	var a, b, c = 4, 3, "hello"

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	// Multiple varsiables 2

	var a1, b1, c1 int = 4, 3, 5 //if type is defined then it must be same

	fmt.Println(a1)
	fmt.Println(b1)
	fmt.Println(c1)

	// Multiple varsiables 3

	a2, b2, c2 := 4, 3, "hello"

	fmt.Println(a2)
	fmt.Println(b2)
	fmt.Println(c2)

	// Multiple varsiables 4

	var (
		a3 int
		b3         = "hello"
		c3 bool    = true
		d3 float32 = 41.15
	)

	fmt.Println(a3)
	fmt.Println(b3)
	fmt.Println(c3)
	fmt.Println(d3)

}
