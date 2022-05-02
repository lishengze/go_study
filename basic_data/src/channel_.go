package main

import (
	"fmt"
	"math/rand"
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

func test_stop_channel() {
	dataCha := make(chan int, 100)
	stopCha := make(chan struct{})

	test_count := 100
	max := 50

	for i := 0; i < test_count; i++ {
		go func(index int) {
			fmt.Println(index)
			randVal := rand.Intn(max)
			select {
			case <-stopCha:
				fmt.Println("Stop Sending Data")
				return
			case dataCha <- randVal:
				fmt.Printf("Send %d To dataCha \n", randVal)
			}
		}(i)
	}

	go func() {
		for value := range dataCha {
			if value == max-1 {
				fmt.Println("Close StopCha")
				close(stopCha)
			}
		}
	}()

	select {
	case <-time.After(time.Hour):
	}
}

func test_close_channel3() {
	senderNumbs := 10
	receiverNumbs := 10
	max := 100

	dataCha := make(chan int, 100)
	sigCha := make(chan struct{})
	modCha := make(chan string)

	fmt.Println("test_close_channel3")

	// Receive goroutine
	for i := 0; i < receiverNumbs; i++ {
		go func(index int) {
			fmt.Printf("[R] Recv Goroutine: %d \n", index)
			select {

			case value := <-dataCha:
				fmt.Printf("[R] Recv Index: %d, Data: %d, \n", index, value)
				if value > max/2 {
					fmt.Printf("[R] Recv Send Close Signal")
					modCha <- "close"
				}
			}
		}(i)
	}

	// moderator goroutine
	go func() {
		for {
			select {
			case recv_value := <-modCha:
				if recv_value == "close" {
					fmt.Printf("\n--------- Close SigCha -------- \n")
					close(sigCha)
				}
				return
			}

		}
	}()

	// Sender goroutine
	for i := 0; i < senderNumbs; i++ {
		go func(index int) {
			fmt.Printf("[S] Send Goroutine: %d \n", index)
			select {
			case <-sigCha:
				fmt.Printf("[S] SigCha Closed, Index: %d, \n", index)
				return
			default:
				value := rand.Intn(max)
				fmt.Printf("[S] Send Index: %d, Data: %d \n", index, value)
				dataCha <- value
			}
		}(i)
	}

}

func write_c(c chan<- int) {
	var i int = 0
	for {
		c <- i

		fmt.Printf("%v, Write %d.\n", time.Now(), i)
		// fmt.Println(time.Now())

		time.Sleep(3 * time.Second)

		i++
	}
}

func read_c(c <-chan int) {
	for {
		i := <-c
		fmt.Printf("%v, Read %d.\n", time.Now(), i)
		// fmt.Println(time.Now())
		time.Sleep(3 * time.Second)

	}
}

func test_read_write() {
	c_value := make(chan int)

	go write_c(c_value)

	go read_c(c_value)

	time.Sleep(1000 * time.Second)

}

func main() {
	// test1()

	// test2()

	// test_select()

	// test_select2()

	// test_write()

	// test_channel2()

	// test_stop_channel()

	// test_close_channel3()

	// select {
	// case <-time.After(time.Hour):
	// }

	test_read_write()
}
