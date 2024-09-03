package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println(s)
	}
}

func sum(s []int, channel chan int) {
	total := 0
	for _, v := range s {
		total += v
	}

	channel <- total
}

func main() {
	go say("world")
	say("hello")

	s := []int{7, 2, 8, -9, 4, 0}
	channel := make(chan int)
	go sum(s[:len(s)/2], channel)
	go sum(s[len(s)/2:], channel)
	x, y := <-channel, <-channel
	fmt.Println(x, y, x+y)
}
