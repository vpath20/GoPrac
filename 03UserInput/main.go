package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	welcome := "Welcome Text"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("EnterYour Messege: ")

	// Comma ok || Error ok

	input, _ := reader.ReadString('\n')

	// ALSO == // input, err := reader.ReadString('\n')   //Don't have try catch in go so we use this comma ok systax

	fmt.Println("Thanks for messege: ", input)
	fmt.Printf("Type of input %T", input)

}
