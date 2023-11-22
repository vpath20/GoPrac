package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	// performPostJSONRequest()
	performPostFormRequest()
}

func performPostJSONRequest() {
	const MyUrl = "http://localhost:3000/post"

	// fake json payload

	requstBody := strings.NewReader(`
	{
		"coursename":"GoLang",
		"price":150,
		"platform":"Youtube"
	}`)

	response, err := http.Post(MyUrl, "application/json", requstBody)

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println("content: ", string(content)) // method 1

	var data strings.Builder // method 2
	data.Write(content)

	fmt.Println("data:", data.String())

	// content, _ := ioutil.ReadAll(response.Body)
	// defer response.Body.Close()
	// fmt.Println("response", response)
	// fmt.Println("content", string(content))

}

func performPostFormRequest() {
	const MyUrl = "http://localhost:3000/post"

	data := url.Values{}
	data.Add("FirstName", "Rohit")
	data.Add("LastName", "Varshney")
	data.Add("Email", "Rohitvarshney@gmail.com")

	response, err := http.PostForm(MyUrl, data)

	if err != nil {
		panic(err)
	}

	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))

}
