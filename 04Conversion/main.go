package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome...")
	fmt.Println("Enter Some Number")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	num, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err == nil {
		fmt.Printf("You Entered: %v we change : %v \n", num, num+2)
	} else {
		fmt.Println(err)
	}
}
