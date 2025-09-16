package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func sum(s []int, c chan int) {
	fmt.Println("=-=-=")
	sum := 0
	for _, v := range s {
		time.Sleep(100 * time.Millisecond)
		fmt.Println("--->", sum)
		sum += v
	}
	c <- sum // send sum to c
}

func main() {
	go say("world")
	say("hello")
	fmt.Println("================")

	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // receive from c

	fmt.Println(x, y, x+y)

}
