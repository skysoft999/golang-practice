package main

import "fmt"

/*
Running the code above will give a deadlock error because the for..range
is waiting to receive a value from the integer channel, but it will never happen.
Also, since there are no other goroutines to execute, the system will raise a deadlock error.
*/
func main() {
	channel := make(chan int)

	go func(channel chan int) {

		for i := 0; i < 100; i++ {
			channel <- i
		}
	}(channel)

	for val := range channel {
		fmt.Println(val)
	}

}

/*
The select function randomly selects one of the case statements to execute
from all the available statements that have data to process.
If no case statement is ready to process, the default statement will be executed.
*/

// package main

// func main() {
// 	c := make(chan int)
// 	select {
// 	case <-c:
// 	}
// }
