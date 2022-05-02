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

type MyContext struct {
	context.Context
}

func work(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("%v, Main Send Done \n", time.Now())
			return
		default:
			time.Sleep(time.Second * 3)
		}
	}
}

func test_mycontext() {
	base_ctx := context.Background()

	base_child_ctx, base_child_cancel_func := context.WithCancel(base_ctx)

	// base_child_ctx

	user_ctx := MyContext{base_ctx}

	user_child_ctx, user_child_cancel_func := context.WithCancel(user_ctx)

	go work(base_child_ctx)

	go work(user_child_ctx)

	base_child_cancel_func()

	user_child_cancel_func()

	fmt.Printf("base_ctx: %v \n", base_ctx)
	fmt.Printf("base_child_ctx: %v \n", base_child_ctx)
	fmt.Printf("user_ctx: %v \n", user_ctx)
	fmt.Printf("user_child_ctx: %v \n", user_child_ctx)

	time.Sleep(time.Hour)

	// user_child_ctx

}

func main() {
	fmt.Println("------ Test Context ------")

	// test_simple()

	test_mycontext()
}
