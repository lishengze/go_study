package BaseData

import (
	"fmt"
	"math"
	"strconv"
)

func test_string_a() {
	data := "Hello Go 语言"
	for index, value := range data {
		fmt.Printf("Index %d is %#U, %x \n", index, value, value)
	}
}

// prec 代表小数位数
func TrunFloatA(f float64, prec int) float64 {
	x := math.Pow10(prec)
	return math.Trunc(f*x) / x
}

func test_trunc_map() {
	f := 2.0 / 3

	// 截断 8 位小数
	f2 := TrunFloatA(f, 8)

	// 输出结果 0.6666666666666666
	fmt.Printf("%v\n", f2)

	// -1 参数表示保持原小数位数，千万要注意，如果你指定了位数就会四舍五入了
	d2 := strconv.FormatFloat(f2, 'f', -1, 64)

	// 输出结果 0.6666666666666666
	fmt.Printf("%s\n", d2)
}

func change_map(m map[int]int) {
	fmt.Printf("%v\n", m)

	m[0] = 100
}

func test_map_func() {
	m := make(map[int]int, 3)
	m[0] = 1

	change_map(m)

	fmt.Printf("m[%d]: %d \n", 0, m[0])
}

type TestStruct struct {
	A int
	B float64
}

func test_map_init() {
	var data map[int]*TestStruct

	data[0].A = 100
	data[0].B = 10.1

	fmt.Println(data)

	var test TestStruct
	test.A += 100
	test.B += 1.1
	fmt.Println(test)

}
