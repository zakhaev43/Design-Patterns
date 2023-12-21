package main

import (
	"fmt"
	"sync"
)

func starter(wg *sync.WaitGroup) {
	fmt.Println("This is the starter on call")
	defer wg.Done()
}

func follow() {
	fmt.Println("This is the follower on call")
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go starter(&wg)
	follow()
	wg.Wait()
}
