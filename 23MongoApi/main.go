package main

import (
	"fmt"
	"net/http"

	"github.com/rohitvarshney/router"
)

func main() {
	fmt.Println("MongoDB API...")
	r := router.Router()
	fmt.Println("Server is getiing started")
	fmt.Println("Listening at port 4000")
	http.ListenAndServe(":4000", r)
}
