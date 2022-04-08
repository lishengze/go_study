package main

import "fmt"

func slice_info() {
	fmt.Printf("")
}

func init_() {
	var slice1 = []int{1, 2, 3, 4}
	var slice2 = make([]int, 5)
	var slice3 = make([]int, 5, 7)
	var slice4 = []int{1, 2, 3, 4}

	fmt.Printf("slice1: %T, slice2: %T, slice3: %T, slice4: %T \n", slice1, slice2, slice3, slice4)

}

func main() {
	init_()
}
