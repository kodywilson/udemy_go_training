package main

import "fmt"

func main() {
	var n, biggest int
	numbers := []int{43, 56, 87, 12, 45, 57, 103, 22, 9}

	for _, v := range numbers {
		if v > n {
			n = v
			biggest = n
		}
	}

	fmt.Println("The biggest number is", biggest)
}
