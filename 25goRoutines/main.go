package main

import (
	"fmt"
	"net/http"
	"sync"
)

// func main() {
// 	go helloworld()
// 	go goodbye()
// 	time.Sleep(2 * time.Second)
// }

// func helloworld() {
// 	fmt.Println("Hello World!")
// }

// func goodbye() {
// 	fmt.Println("Good Bye!")
// }

// ===== instead of sleep we can use sync.WaitGroup=====

// func main() {
// 	var wg sync.WaitGroup
// 	wg.Add(2)
// 	go helloworld(&wg)
// 	go goodbye(&wg)
// 	wg.Wait()

// }

// func helloworld(wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	fmt.Println("Hello World!")
// }

//	func goodbye(wg *sync.WaitGroup) {
//		defer wg.Done()
//		fmt.Println("Good Bye!")
//	}
var wg sync.WaitGroup //pointer
var signals = []string{"test"}
var mut sync.Mutex //pointer

func main() {
	websiteList := []string{
		"https://google.com",
		"https://go.dev",
		"https://fb.com",
		"https://leetcode.com",
		"https://fb.com",
		"https://github.com",
	}

	for _, web := range websiteList {
		go getStatusCode(web)
		wg.Add(1)
		go getStatusCode(web)
		wg.Add(1)

	}
	wg.Wait()
	fmt.Println(signals)

}

func getStatusCode(endpoint string) {
	defer wg.Done()
	res, err := http.Get(endpoint)
	if err != nil {
		fmt.Println("OOPS in endpoint")
	} else {
		mut.Lock()
		signals = append(signals, endpoint)
		mut.Unlock()
		fmt.Printf("%d Status code for website %s \n", res.StatusCode, endpoint)

	}
}
