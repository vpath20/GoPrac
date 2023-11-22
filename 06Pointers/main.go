package main

import "fmt"

func main() {

	// var ptr *int
	// fmt.Println(ptr)

	myNumber := 45
	var ptr = &myNumber

	fmt.Println("value of actual pointe:r", ptr)
	fmt.Println("value of actual pointe:r", *ptr)

	*ptr = *ptr * 2
	fmt.Println(myNumber)
	fmt.Println(*ptr)

	//    pointers makes it a guarantee that no matter what you are doing the operation with those valllues those operation are actually perform on the values not on the copies of those values

}
