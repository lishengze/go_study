package BaseData

import (
	"fmt"
	"math"
	"math/big"
)

func parse_float() {
	var f_data float32 = 0.085
	bits := math.Float32bits(f_data) // trans_api
	bit_array := fmt.Sprintf("%.32b", bits)
	fmt.Printf("%s | %s | %s \n", bit_array[0:1], bit_array[1:9], bit_array[9:])
}

// 问题:
// 1. 数组参数的声明，切片参数的声明;
//

func trans_bit_to_float32() {
	var f_data float32 = 0.085
	bits := math.Float32bits(f_data)
	bit_array := fmt.Sprintf("%.32b", bits)

	fmt.Printf("bits: %b, bit_array: %s\n", bits, bit_array)

	sign := bits & (1 << 31)

	exponent_raw := int(bits >> 23)
	exponent_bias := 127
	exponent := exponent_raw - exponent_bias

	// compute decimal fractional part
	var fraction float64 = 0

	for index, bit := range bit_array[9:32] {
		fmt.Printf("index:%d, bit:%d \n", index, bit)

		if bit == 49 {
			position := index + 1
			curr_value := math.Pow(2, float64(position))
			fraction += 1 / curr_value
		}

	}

	result := (1 + fraction) * math.Pow(2, float64(exponent))

	fmt.Printf("sign: %d, exponent: %d(%d), fraction: %f, result: %f \n\n",
		sign, exponent, exponent_raw, fraction, result)
}

func is_int() {
	var f_data float32 = 2345.1
	bits := math.Float32bits(f_data)
	exponent := int(bits>>23) - 127 - 23

	if exponent >= 0 {
		fmt.Printf("exponent: %d, %10f, is integer", exponent, f_data)
		return
	}

	fractional := bits&((1<<23)-1) | (1 << 23)
	result := fractional & (1<<(-exponent) - 1)

	fmt.Printf("exponent: %d, fractional: %b, result: %b\n", exponent, fractional, result)

	if result != 0 {
		fmt.Printf("%.10f is not integer", f_data)
	} else {
		fmt.Printf("%.10f is integer", f_data)
	}
}

func test_big() {
	var x1, y1 float64 = 10, 3
	z1 := x1 / y1
	fmt.Println(x1, y1, z1)

	x2, y2 := big.NewFloat(10), big.NewFloat(3)
	x2.SetPrec(100)
	y2.SetPrec(100)
	z2 := new(big.Float).Quo(x2, y2)
	fmt.Println(x2, y2, z2)
}

// func main() {
// 	// parse_float()

// 	// trans_bit_to_float32()

// 	// is_int()

// 	test_big()
// }
