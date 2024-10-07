package main

import (
	"time"
)

// This is an example of race condition
// 2 goroutines try to write and there is no access control.

var x int = 0

func decrement() {
	for {
		x = x - 50
	}
}

func increment() {
	for {
		x = x + 15
	}
}

func main() {
	go increment()
	go decrement()
	time.Sleep(5 * time.Second)
}
