package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func main() {
	url := "https://randomuser.me/api/?page=3&results=10&seed=abc"

	response, err := http.Get(url)

	if err != nil {
		fmt.Println("Error Getting response : ", err)
	}

	defer response.Body.Close()
	fmt.Println("response status", response.Status)
	fmt.Println("response content legth", response.ContentLength)

	body, err := io.ReadAll(response.Body)

	if err != nil {
		fmt.Println("Error while reading : ", err)
	}

	var jsonReadable bytes.Buffer // beautiful json human readable
	err = json.Indent(&jsonReadable, body, "", " ")
	if err != nil {
		fmt.Println(err)
	}

	// fmt.Println("content", string(body))

	// var responseString strings.Builder //extra method 2
	// byteCount, _ := responseString.Write(body)
	// fmt.Println("byteCount", byteCount)
	// fmt.Println("data:", responseString.String())

	fmt.Println("Response Content:")
	fmt.Println(jsonReadable.String())

}
