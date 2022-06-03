package BaseData

import "fmt"

func test_array() {
	a := []int{1, 2, 3, 4}

	for i := 0; i < 10; i++ {
		fmt.Println(a[i])
	}

	data_map := map[string]int{
		"a": 1,
	}

	for k, v := range data_map {
		fmt.Println("%s, %d", k, v)
	}

	// fmt.Println(a[4])
}

// func main() {
// 	test_array()

// 	fmt.Println("Test Array")
// }
