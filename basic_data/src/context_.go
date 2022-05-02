package main

import (
	"context"
	"fmt"
	"time"
)

func gen(ctx context.Context) chan int {

	c_i := make(chan int)

	go func() {
		i := 0
		for {

			select {
			case <-ctx.Done():
				fmt.Printf("Gen Receive Done!\n")
				return
			case c_i <- i:
				fmt.Printf("Write %d\n", i)
				i++
				time.Sleep(time.Second * 3)

			}
		}
	}()

	fmt.Println("gen return")
	return c_i
}

func test_simple() {
	ctx, cancel_func := context.WithCancel(context.Background())

	defer cancel_func()

	for i := range gen(ctx) {
		fmt.Printf("Read %d\n", i)

		if i == 5 {
			fmt.Println("Main Send Done")
			cancel_func()
			break
		}
	}
}

func main() {
	fmt.Println("------ Test Context ------")

	test_simple()
}
