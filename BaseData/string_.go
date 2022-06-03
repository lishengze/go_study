package BaseData

import (
	"fmt"
	"math"
	"strconv"
)

func test_string() {
	data := "Hello Go 语言"
	for index, value := range data {
		fmt.Printf("Index %d is %#U, %x \n", index, value, value)
	}
}

// prec 代表小数位数
func TrunFloat(f float64, prec int) float64 {
	x := math.Pow10(prec)
	return math.Trunc(f*x) / x
}

func test_trunc() {
	f := 2.0 / 3

	// 截断 8 位小数
	f2 := TrunFloat(f, 8)

	// 输出结果 0.6666666666666666
	fmt.Printf("%v\n", f2)

	// -1 参数表示保持原小数位数，千万要注意，如果你指定了位数就会四舍五入了
	d2 := strconv.FormatFloat(f2, 'f', -1, 64)

	// 输出结果 0.6666666666666666
	fmt.Printf("%s\n", d2)
}

// func main() {
// 	// test_string()

// 	test_trunc()
// }
