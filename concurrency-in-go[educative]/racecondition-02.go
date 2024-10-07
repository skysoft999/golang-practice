package main

import (
	"time"
)

/*
You’ll get the error fatal error: concurrent map read and map write
if you execute the above code widget because read and write
 goroutines are accessing the sharedMap variable at the same time.
*/

func read(sharedMap map[int]int) {
	for {
		var _ int = sharedMap[0]
	}
}

func write(sharedMap map[int]int) {
	for {
		sharedMap[0] = sharedMap[0] + 1
	}
}

func main() {
	sharedMap := make(map[int]int)
	sharedMap[0] = 0

	go read(sharedMap)
	go write(sharedMap)
	time.Sleep(2 * time.Second)
}
