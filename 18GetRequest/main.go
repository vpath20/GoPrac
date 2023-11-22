package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	GetRequest()
}

func GetRequest() {
	const myUrl = "http://localhost:3000/get"

	response, err := http.Get(myUrl)

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	fmt.Println("response status", response.Status)
	fmt.Println("response content legth", response.ContentLength)

	content, _ := ioutil.ReadAll(response.Body) // method 1
	fmt.Println("content", string(content))     //method 1

	var responseString strings.Builder //extra method 2
	byteCount, _ := responseString.Write(content)
	fmt.Println("byteCount", byteCount)
	fmt.Println("data:", responseString.String())

}
