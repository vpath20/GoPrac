package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const URL = "https://lco.dev"

func main() {
	response, err := http.Get(URL)

	if err != nil {
		panic(err)
	}
	fmt.Printf("response %T  === %v", response, response)

	defer response.Body.Close()

	dataBytes, err2 := ioutil.ReadAll(response.Body)

	if err2 != nil {
		panic(err2)
	}
	fmt.Printf(string(dataBytes))

}
