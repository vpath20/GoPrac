package main

import "fmt"

func main() {
	languages := make(map[string]string)

	languages["JS"] = "javascript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Pyhton"
	fmt.Println(languages)
	fmt.Println(languages["JS"])

	delete(languages, "RB")
	fmt.Println(languages)

	for key, value := range languages {
		fmt.Println("key", key, " value", value)
	}
}
