package main

import "fmt"

/*
测试函数作为 参数， 返回值
*/

func add(a int, b int) int {
	fmt.Printf("Add a: %d, b: %d \n", a, b)
	return a + b
}

func minus(a int, b int) int {

	fmt.Printf("minus a: %d, b: %d \n", a, b)

	return a - b
}

func test(a int, b int, before_func func(int, int) int) (after_func func(int, int) int) {
	fmt.Printf("test a: %d, b: %d , \n", a, b)

	fmt.Printf("before_func(a,b): %d \n", before_func(a, b))

	return minus
}

func main() {
	a := 10
	b := 20
	fmt.Printf("test(a, b, add)(a, b): %d \n", test(a, b, add)(a, b))
}
