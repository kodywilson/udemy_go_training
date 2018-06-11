package main

import "fmt"

func main() {

	var small int
	var large int

	fmt.Println("Please enter a small number:")
	fmt.Scan(&small)
	fmt.Println("Please enter a large number:")
	fmt.Scan(&large)
	fmt.Println(large, "divided by", small, " = ", (large / small))

}
