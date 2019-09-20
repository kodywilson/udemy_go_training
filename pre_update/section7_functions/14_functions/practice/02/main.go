package main

import "fmt"

// take integers and add them all up
func calc_add(x ...float64) float64 {
	total := 0.0
	for _, v := range x {
		total += v
	}
	return total
}

func main() {
	numbers := []float64{10, 20, 30}
	added := calc_add(numbers...)
	fmt.Println(added)
}
