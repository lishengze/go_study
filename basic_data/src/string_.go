package main

import "fmt"

func test_string() {
	data := "Hello Go 语言"
	for index, value := range data {
		fmt.Printf("Index %d is %#U, %x \n", index, value, value)
	}
}

func main() {
	test_string()
}
