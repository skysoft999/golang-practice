package main

import "fmt"

func countStr(s string) map[string]int {
	countMap := make(map[string]int)
	for _, v := range s {
		char := string(v)
		countMap[char]++
	}
	return countMap

}

func main() {
	input_str := "elephant"
	fmt.Println(countStr(input_str))
}
