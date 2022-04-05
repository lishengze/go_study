package main

import (
	"fmt"
	"sync"
)

var lock = &sync.Mutex{}

type single struct {
}

var g_single *single

func GetSingle() *single {

	if g_single == nil {
		lock.Lock()
		defer lock.Unlock()

		if g_single == nil {
			g_single = new(single)
			fmt.Println("Init Single")
		} else {
			fmt.Println("Second Judge")
		}
	} else {
		fmt.Println("Single already created!")
	}

	return g_single
}

func main() {
	fmt.Println("Test Signaton")

	for i := 0; i < 20; i++ {
		go GetSingle()
	}

	fmt.Scanln()
}
