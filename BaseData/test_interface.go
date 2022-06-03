package BaseData

import (
	"fmt"
	"unsafe"
)

type Shape interface {
	area() float64
	perimeter() float64
}

type Rectangle struct {
	a, b float64
}

func (r Rectangle) area() float64 {
	return r.a * r.b
}

func (r Rectangle) perimeter() float64 {
	return (r.a + r.b) * 2
}

func test_shape() bool {
	return true
}

type Duck interface {
	Quark()
}

type Cat struct {
}

// func (c Cat) Quark() {
// 	fmt.Println("Cat Print")
// }

func (c *Cat) Quark() {
	fmt.Println("Cat* Quark")
}

func test_cat_pointer() bool {

	// var duck Duck = Cat{}
	// duck.Quark()

	return true
}

type ReaderUse interface {
	Read()
}

type WriterUse interface {
	Write()
}

// func test_io() {
// 	var r  io.Reader()
// 	tty, err := os.Open("test.txt", os.O_RDWR|os.O_CREATE)
// 	if err == nil {
// 		fmpt.Println("Open Failed")
// 	}

// 	r = tty

// 	var w io.Writer
// 	w = r.(io.Writer)
// }

type Coder interface {
	Code()
}

type Gopher struct {
	Name string
}

func (g *Gopher) Code() string {
	return g.Name
}

func test_code() {
	var code Coder
	fmt.Println(code == nil)
	fmt.Printf("code: %v, %T\n", code, code)

	var gopher *Gopher
	fmt.Println(gopher == nil)
	fmt.Printf("gopher: %v, %T\n", gopher, gopher)

	// code = gopher
	// fmt.Println(code == nil)
	// fmt.Printf("code: %v, %T", code, code)

}

func test_itab() {
	type iface struct {
		itab, data uintptr
	}

	var x interface{} = nil
	var y interface{} = (*int)(nil)
	var data int = 5
	var z interface{} = &data

	ix := *(*iface)(unsafe.Pointer(&x))
	iy := *(*iface)(unsafe.Pointer(&y))
	iz := *(*iface)(unsafe.Pointer(&z))

	fmt.Println(ix.itab, ix.data)
	fmt.Println(iy.itab, iy.data)
	fmt.Println(iz.itab, iz.data)

	fmt.Println(*(*int)(unsafe.Pointer(iz.data)))
}

// func main() {

// 	fmt.Println("--------- test interface ---------")

// 	// test_code()

// 	test_itab()

// 	// test_cat_pointer()
// }
