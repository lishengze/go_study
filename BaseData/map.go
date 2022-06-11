package BaseData

import (
	"fmt"
	"sync"
)

var a sync.Map

func output(key, value interface{}) bool {
	fmt.Printf("key:%s, value: %d \n", key, value)
	if value != 0 {
		a.Store(key, 0)
	}
	return true
}

func test_sync_map() {
	// var a sync.Map
	a.Store("a", 1)
	a.Store("b", 2)

	a.Range(output)

	a.Range(output)
}

func TestMapMain() {
	test_sync_map()
}
