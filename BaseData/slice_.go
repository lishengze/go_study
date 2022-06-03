package BaseData

import (
	"fmt"
	"unsafe"
)

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

func TestSlice() {
	a := make([]int, 0, 5)

	b := make([]int, 3)

	fmt.Printf("a: %+v \n", a)

	fmt.Printf("b: %+v \n", b)
}

func TestCopy() {
	// a := []int{1, 2, 3}

	// copy(b, a)

	a := []int{1, 2, 3}
	// b := []int{-1, -4}
	b := []int{-1}
	// b := make([]int, 0)
	fmt.Printf("a: %+v, len(a): %d, cap(a): %d \n", a, len(a), cap(a))
	fmt.Printf("b: %+v len(b): %d, cap(b): %d \n", b, len(b), cap(b))

	copy(b, a)
	fmt.Println(unsafe.Pointer(&a)) // 0xc0000a4018
	fmt.Println(a, &a[0])           // [1 2 3] 0xc0000b4000
	fmt.Println(unsafe.Pointer(&b)) // 0xc0000a4030
	fmt.Println(b, &b[0])

	// fmt.Printf("a: %+v \n", a)
	// fmt.Printf("b: %+v \n", b)

}

func TestSliceMain() {
	// TestSlice()

	TestCopy()
}

// func main() {
// 	init_()
// }
