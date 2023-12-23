package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Channels in golang")

	myCh := make(chan int, 2)
	wg := &sync.WaitGroup{}

	// fmt.Println(<-myCh)
	// myCh <- 5

	wg.Add(2)
	// ===== Recieve ONLY
	go func(ch <-chan int, wg *sync.WaitGroup) {
		val, isChanOpen := <-ch

		fmt.Println(val)
		fmt.Println(isChanOpen)

		// fmt.Println(<-ch)
		wg.Done()
	}(myCh, wg)

	//===== Send ONLY
	go func(ch chan<- int, wg *sync.WaitGroup) {
		// ch <- 5
		// ch <- 6
		close(ch)
		// ch <- 0  //gives error because sending value after closing the channel
		wg.Done()

	}(myCh, wg)
	wg.Wait()

}
