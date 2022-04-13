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

func test_element() {
	a := 100
	b := &a
	var c interface{} = b

	r_c := reflect.ValueOf(&c)

	fmt.Println(r_c.Kind())
	fmt.Println(r_c.CanSet())
	fmt.Println("")

	fmt.Println(r_c.Elem().Kind())
	fmt.Println(r_c.Elem().CanSet())
	fmt.Println("")

	fmt.Println(r_c.Elem().Elem().Kind())
	fmt.Println(r_c.Elem().Elem().CanSet())
	fmt.Println("")

	fmt.Println(r_c.Elem().Elem().Elem().Kind())
	// fmt.Println((r_c.Elem().Elem().Elem().CanSet())
	fmt.Println("")
}

func get_reflect_info(src reflect.Value) {
	fmt.Println(src.Kind())

	switch src.Kind() {
	case reflect.Int:
		fmt.Println(src.Int())
	case reflect.String:
		fmt.Println(src.String())
	case reflect.Float32:
		fmt.Println(src.Float())
	case reflect.Float64:
		fmt.Println(src.Float())
	}
}

func base_test() {
	a := 10
	b := "b"

	f_a := reflect.ValueOf(a)
	// fmt.Println(f_a.Kind())

	f_b := reflect.ValueOf(b)
	// fmt.Println(f_b.Kind())

	get_reflect_info(f_a)
	get_reflect_info(f_b)

}

type User struct {
	Name  string
	Age   int
	Grade int
}

type data interface {
}

// Type 获取key 名;
// Value 根据 key, 获取value;
func test_value() {
	user := User{"Tom", 10, 1}

	r_u := reflect.ValueOf(&user)
	t_u := reflect.TypeOf(&user)
	fmt.Printf("r_u: %v \n", t_u.Kind())

	fmt.Println(r_u)

	for i := 0; i < t_u.Elem().NumField(); i++ {
		// fmt.Printf("value.key %v, value.type %v, value: %v, canset: %v \n", t_u.Field(i).Name, t_u.Field(i).Type, r_u.Field(i), r_u.Field(i).CanSet())
		fmt.Printf("value.key %v, value.type %v, value: %v, canset: %v \n",
			t_u.Elem().Field(i).Name, t_u.Elem().Field(i).Type,
			r_u.Elem().Field(i), r_u.Elem().Field(i).CanSet())
	}
}

func main() {
	// test_reflect1()

	// base_test()

	// test_element()

	test_value()
}
