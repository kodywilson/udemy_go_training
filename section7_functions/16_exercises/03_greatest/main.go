package main

import "fmt"

func greatest(numbers ...int) int {
	var n, biggest int
	for _, v := range numbers {
		if v > n {
			n = v
			biggest = n
		}
	}
	return biggest
}

func main() {
	greatest := greatest(43, 56, 87, 12, 45, 57, 103, 22, 9)
	fmt.Println("The biggest number is", greatest)
}
