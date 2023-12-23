package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Race Condition ")

	wg := &sync.WaitGroup{}
	mut := &sync.RWMutex{}
	var score = []int{0}

	wg.Add(4)

	go func(wg *sync.WaitGroup, mut *sync.RWMutex) {
		fmt.Println("First R")
		mut.Lock()
		score = append(score, 1)
		mut.Unlock()
		wg.Done()
	}(wg, mut)

	go func(wg *sync.WaitGroup, mut *sync.RWMutex) {
		fmt.Println("Second R")
		mut.Lock()
		score = append(score, 2)
		mut.Unlock()
		wg.Done()
	}(wg, mut)
	go func(wg *sync.WaitGroup, mut *sync.RWMutex) {
		fmt.Println("Third R")
		mut.Lock()
		score = append(score, 3)
		mut.Unlock()
		wg.Done()
	}(wg, mut)

	go func(wg *sync.WaitGroup, mut *sync.RWMutex) {
		fmt.Println("Third Read")
		mut.RLock()
		fmt.Println(score)
		mut.RUnlock()
		wg.Done()
	}(wg, mut)

	wg.Wait()
	fmt.Println(score)

}
