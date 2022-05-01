package main

import (
	"fmt"
	"runtime"
	"time"
)

func test_trap() {
	threads_count := runtime.GOMAXPROCS(0) + 10
	fmt.Printf("threads_count: %d\n", threads_count)
	x := 0
	for i := 0; i < threads_count; i++ {
		go func() {
			for {
				x++
			}
		}()
	}
	time.Sleep(time.Second)
	fmt.Printf("x= %d\n", x)
}

func test_trap2() {
	var x int
	threads := runtime.GOMAXPROCS(0)
	println(threads)
	for i := 0; i < threads; i++ {
		go func() {
			for {
				x++
			}
		}()
	}
	time.Sleep(time.Second)
	fmt.Println("x =", x)
}

func main() {
	// test_trap()

	test_trap2()
}
