package main

import (
	"fmt"
	"sync"
)

func main() {
	mu1 := sync.Mutex{}
	mu2 := sync.Mutex{}

	mu1.Lock()
	mu2.Lock()
	fmt.Println("in mutex")
	mu2.Unlock()
	mu1.Unlock()

	fmt.Println("end mutex")
}
