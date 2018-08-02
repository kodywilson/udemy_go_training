package main

import "fmt"

func greatest(numbers ...int) int {
	var biggest int
	for _, v := range numbers {
		if v > biggest {
			biggest = v
		}
	}
	return biggest
}

func main() {
	greatest := greatest(43, 56, 87, 12, 45, 57, 103, 22, 9)
	fmt.Println("The biggest number is", greatest)
}
