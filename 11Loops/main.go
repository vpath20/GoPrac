package main

import "fmt"

func main() {
	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
	fmt.Println(days)

	// for d := 0; d < len(days); d++ {
	// 	fmt.Println(days[d])
	// }

	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	// ===== Foreach =====
	for i, d := range days {
		fmt.Printf("Index : %v value : %v\n", i, d)
	}

	// ===== While Loop =====

	value := 1
	for value < 10 {
		if value == 2 {
			goto kw
		}
		fmt.Println(value)
		value++
	}
	// keywords for goto statement
kw:
	fmt.Println("Jumped at this keyword")

}
