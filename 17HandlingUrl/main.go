package main

import (
	"fmt"
	"net/url"
)

const myUrl = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=ghbj450ghb"

func main() {
	fmt.Println("Handling URL")
	// Parsing the URl

	result, _ := url.Parse(myUrl)

	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	qParams := result.Query()
	fmt.Println(qParams)

	fmt.Println(qParams["coursename"])

	for key, val := range qParams {
		fmt.Println("key is", key)
		fmt.Println("value is", val)
	}

	partsOfUrl := &url.URL{
		Scheme: "https",
		Host:   "lco.dev",
		Path:   "/tutcss",
	}

	anotherUrl := partsOfUrl.String()
	fmt.Println(anotherUrl)

}
