package main

import (
	"fmt"
	"sort"
)

func main() {
	var fruit = []string{"apple", "banana"}
	fmt.Println(fruit)

	fruit = append(fruit, "guava", "mango")
	fmt.Println(fruit)

	fruit = append(fruit, fruit[1:]...)
	fmt.Println(fruit)

	fruit = append(fruit, fruit[:1]...)
	fmt.Println(fruit)

	highScore := make([]int, 4)
	highScore[0] = 234
	highScore[1] = 244
	highScore[2] = 233
	highScore[3] = 434
	// highScore[4] = 45    //error index out of Bound

	highScore = append(highScore, 45, 456)

	fmt.Println(highScore)

	fmt.Println(sort.IntsAreSorted(highScore))
	sort.Ints(highScore)
	fmt.Println(highScore)

	// remove the values in slices

	var courses = []string{"java", "python", "JS", "go"}
	fmt.Println(courses)
	index := 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)

}
