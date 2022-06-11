package BaseData

import (
	"fmt"

	"github.com/emirpasic/gods/maps/treemap"
)

func TestEqualKey() {

	test := treemap.NewWithIntComparator()

	test.Put(1, 10)
	test.Put(2, 11)
	test.Put(3, 12)
	test.Put(3, 13)

	iter := test.Iterator()
	iter.Begin()

	for iter.Next() {
		fmt.Printf("key: %+v , value: %+v\n", iter.Key(), iter.Value())
	}
}

func TestTreeMap() {
	fmt.Println("------- TestTreeMap  -------")

	TestEqualKey()
}

// func main() {
// 	fmt.println("------ Study TreeMap -------")
// }
