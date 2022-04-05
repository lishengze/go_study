package main

import (
	"fmt"
	"reflect"
)

type Child struct {
	Name     string
	Age      int
	Handsome bool
}

type Adult struct {
	ID         string `qson:"Name"`
	Occupation string
	Handsome   bool
}

func HandsomeF(src interface{}) {
	re_value := reflect.ValueOf(src)

	if re_value.Kind() == reflect.Slice {
		fmt.Println("reflect is Slice")
	} else {
		fmt.Println("reflect is not Slice")
	}

	element_value := re_value.Type().Elem()

	if element_value.Kind() == reflect.Struct {
		fmt.Println("element is Struct")
	}

	for i := 0; i < re_value.Len(); i++ {
		cur_value := re_value.Index(i)
		handsome := cur_value.FieldByName("Handsome")

		// var name reflect.Value

		for j := 0; j < element_value.NumField(); j++ {
			field := element_value.Field(j)
			tag := field.Tag.Get("qson")

			if field.Name == "Name" || tag == "name" {
				handsome.SetBool(true)
			}

		}

		// handsome.SetBool(true)
	}

}

func test_reflect1() {
	child_slice := []Child{
		{Name: "yu", Age: 4, Handsome: true},
		{Name: "gu", Age: 5, Handsome: false},
	}

	adult_slice := []Adult{
		{ID: "A", Occupation: "first", Handsome: true},
		{ID: "B", Occupation: "second", Handsome: false},
	}

	fmt.Printf("Adults Before Handsome: %v\n", adult_slice)
	HandsomeF(adult_slice)
	fmt.Printf("Adults After Handsome: %v\n", adult_slice)

	fmt.Printf("child_slice Before Handsome: %v\n", child_slice)
	HandsomeF(child_slice)
	fmt.Printf("child_slice After Handsome: %v\n", child_slice)

}

func main() {
	test_reflect1()
}
