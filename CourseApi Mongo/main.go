package main

import (
	"fmt"
	"net/http"

	"github.com/rohit/courseApi/routers"
)

func main() {
	fmt.Println("API...")
	r := routers.Routers()

	fmt.Println("Server is Getting Started")
	fmt.Println("Listening to port 4000...")
	http.ListenAndServe(":4000", r)
}
