package main

import (
	"fmt"
	"strconv"
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

func test_float() {
	f := 100.12345678901234567890123456789
	fmt.Println(strconv.FormatFloat(f, 'b', 5, 32))
	// 13123382p-17
	fmt.Println(strconv.FormatFloat(f, 'e', 5, 32))
	// 1.00123e+02
	fmt.Println(strconv.FormatFloat(f, 'E', 5, 32))
	// 1.00123E+02
	fmt.Println(strconv.FormatFloat(f, 'f', 5, 32))
	// 100.12346
	fmt.Println(strconv.FormatFloat(f, 'g', 5, 32))
	// 100.12
	fmt.Println(strconv.FormatFloat(f, 'G', 5, 32))
	// 100.12
	fmt.Println(strconv.FormatFloat(f, 'b', 30, 32))

}

func test_go_lang() {

}

func main() {
	fmt.Println("Test Signaton")

	for i := 0; i < 20; i++ {
		go GetSingle()
	}

	fmt.Scanln()
}
