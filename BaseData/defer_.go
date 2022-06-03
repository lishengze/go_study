package BaseData

import "fmt"

// func test1() {
// 	defer fmt.Println("test1 This defer")

// 	panic("Panic By user")

// 	fmt.Println("test1 over!")
// }

func catch_exp() {
	errMsg := recover()

	if errMsg != nil {
		fmt.Println(errMsg)
	}

	fmt.Println("This is catch_exp func")

}

func test2() {
	defer catch_exp()

	fmt.Println("test1 over!")
}

func test3() {
	for i := 1; i < 5; i++ {
		defer func(data int) {
			fmt.Println(data)
		}(i)
	}
}

var g int = 100

func test4() int {

	defer func() {
		g = 200
		fmt.Printf("test4 defer g: %d\n", g)
	}()

	return g
}

func test5() int {
	r := g
	defer func() {
		r = 200
		fmt.Printf("test5 defer r: %d\n", r)
	}()

	r = 0
	return r

}

// func main() {
// 	// test1()

// 	// test2()

// 	// test3()

// 	// fmt.Printf("test4 over g %d\n", test4())

// 	fmt.Printf("test5 over r %d\n", test5())

// 	fmt.Println("main over!")
// }
