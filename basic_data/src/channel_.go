package main

import (
	"fmt"
	"net/http"
	"time"
)

func test1() {
	var c = make(chan int)

	go func() {
		c <- 100
		fmt.Printf("Write c: %d\n", c)
	}()

	go func() {
		data, ok := <-c
		fmt.Println("Read c: ", data, ok)
	}()

	close(c)

	time.Sleep(1 * time.Second)
}

func test2() {
	test_links := []string{
		"www.baidu.com",
		"www.jd.com",
	}

	var c_check_result = make(chan string)
	for _, link := range test_links {

		go check_link(link, c_check_result)
	}

	for l := range c_check_result {
		data, ok := <-c_check_result
		fmt.Println("c_check_result: ", data, ok)
		time.Sleep(1 * time.Second)
		fmt.Println("l: ", l)
	}

	// data, ok := <-c_check_result
	// fmt.Println("c_check_result: ", data, ok)
	// time.Sleep(3 * time.Second)
}

func check_link(link string, c_check_result chan string) {
	_, ok := http.Get(link)

	if ok != nil {
		c_check_result <- link + " is up"
	}
}

func test_select() {
	tick := time.Tick(time.Second)
	for {
		select {
		case <-tick:
			fmt.Println(time.Now())
		}
	}
}

func select_data(a chan int, b chan int) {
	for i := 0; i < 2; i++ {
		select {
		case a <- 1:
			fmt.Println("channel a")
			a = nil
		case b <- 2:
			fmt.Println("channel b")
			b = nil
		}
	}
}

func test_select2() {
	a := make(chan int)
	b := make(chan int)

	go select_data(a, b)

	fmt.Println(<-a)
	fmt.Println(<-b)
}

func grA(c chan int) {
	fmt.Printf("grA data: %d\n", <-c)
}

func grB(c chan int) {
	fmt.Printf("grB data: %d\n", <-c)
}

func test_write() {
	c := make(chan int)

	go grA(c)

	go grB(c)

	c <- 3

	// c <- 4

	time.Sleep(3 * time.Second)
}

type User struct {
	Name string
	Age  int
}

var user = User{"Tom", 14}
var pg_user = &user

func modify_user(user *User) {
	user.Name = "Json"
	fmt.Println("Modified User: ", user)
}

func print_user(channel_user chan *User) {
	fmt.Println("Start Recv Data From Channel ", time.Now())

	channel_value := <-channel_user

	fmt.Println("Recv Data User: ", channel_value, time.Now())
}

// Test Channel Value Copy
func test_channel2() {
	user_channel := make(chan *User, 5)

	fmt.Println("Global User: ", pg_user)

	go print_user(user_channel)

	time.Sleep(3 * time.Second)

	user_channel <- pg_user
	// fmt.Println("Send Channel Data Over")

	go modify_user(pg_user)

	time.Sleep(3 * time.Second)
}

func main() {
	// test1()

	// test2()

	// test_select()

	// test_select2()

	// test_write()

	test_channel2()
}
