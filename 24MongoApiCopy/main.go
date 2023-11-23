package main

import (
	"fmt"

	"net/http"

	"github.com/rohit/controller/router"
)

func main() {
	fmt.Println("API...")
	r := router.Routers()

	fmt.Println("Server is geting started...")
	fmt.Println("Listening to Port 3000")
	http.ListenAndServe(":3000", r)

}
