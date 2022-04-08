package main

import "fmt"

// 数组参数声明 类型 + 个数
func copy(c []int) {

	fmt.Printf("Copyed Data %T, %p\n", c, &c)

}

func test() {
	a := [...]int{1, 2, 3}
	fmt.Printf("Original Data %T, %p\n", a, &a)

	copy(a[:])
}

func main() {
	test()
}
