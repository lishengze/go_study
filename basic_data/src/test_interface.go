package main

import (
	"fmt"
	"io"
	"os"
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

	var duck Duck = Cat{}
	duck.Quark()

	return true
}

type ReaderUse interface {
	Read()
}

type WriterUse interface {
	Write()
}

func test_io() {
	var r  io.Reader()
	tty, err := os.Open("test.txt", os.O_RDWR|os.O_CREATE)
	if err == nil {
		fmpt.Println("Open Failed")
	}

	r = tty

	var w io.Writer
	w = r.(io.Writer)
}

func main() {
	fmt.Println("test interface")

	test_cat_pointer()
}
