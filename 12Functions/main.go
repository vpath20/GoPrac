package main

import (
	"fmt"
)

func main() {
	fmt.Println("welcome")
	greeting()
	result := adder(1, 5)
	fmt.Println("result = ", result)

	result2, mssg := proAdder(1, 3, 41, 45, 31)
	fmt.Println("result2 = ", result2, mssg)

	anonymus()

	counter1 := counter()
	fmt.Println(counter1())
	fmt.Println(counter1())
	fmt.Println(counter1())
	fmt.Println(counter1())

	counter2 := counter()
	fmt.Println(counter2())
	fmt.Println(counter2())
	fmt.Println(counter2())
	fmt.Println(counter2())

	fmt.Println("===")
	fib := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(fib())
	}

	// fac := factorial(5)
	// fmt.Println("Factorial of number: ", fac)

}

func greeting() {
	fmt.Println("Hii! Hello!")
}

func adder(num1 int, num2 int) (r int) {
	r = num1 + num2
	return
}

func proAdder(values ...int) (int, string) {
	total := 0
	for _, val := range values {
		total += val
	}
	return total, "This one is a pro adder"
}

func anonymus() {
	fmt.Println("This is function")

	func() {
		fmt.Println("This is called as anonymous function which is called during initialization ")
	}()
	a := func() {
		fmt.Println("This is another way of anonymous function which can called like this ")
	}
	a()

}

// Closure function
func counter() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return b - a

	}

}

// func factorial(n int) func(n int) int {
// 	a, b := 1, 1
// 	return func(n int) int {
// 		if n > 1 {
// 			a, b = b, a*b
// 		} else {
// 			return 1
// 		}
// 		return b
// 	}

// }

// package main

// import "fmt"

// func Concat(sep string, tokens ...string) (r string) {
// 	for i, t := range tokens {
// 		if i != 0 {
// 			r += sep
// 		}
// 		r += t
// 	}
// 	return
// }

// func main() {
// 	tokens := []string{"Go", "C", "Rust"}
// 	// manner 1
// 	langsA := Concat(",", tokens...)
// 	// manner 2
// 	langsB := Concat(",", "Go", "C", "Rust")
// 	fmt.Println(langsA == langsB) // true
// }
