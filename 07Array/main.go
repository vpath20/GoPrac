package main

import "fmt"

func main() {
	var fruit [4]string
	fruit[0] = "apple"
	fruit[1] = "banana"
	fruit[3] = "guava"

	fmt.Println(fruit)

	vegetables := [3]string{"potato", "tomato", "beans"}
	fmt.Println(vegetables)

}
