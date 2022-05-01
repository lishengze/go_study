package main

import (
	"fmt"
	"time"
)

func TimeTask(interval int64) {
	fmt.Println("TimeTask")
	duration := time.Duration(3 * time.Second)
	timer := time.Tick(duration)
	for {
		select {
		case <-timer:
			fmt.Println(time.Now())
		}
	}
}

func main() {
	go TimeTask(5)

	time.Sleep(time.Second * 1000)
}
