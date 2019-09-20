package main

import "fmt"

func main() {
	mySlice := make([]int, 1)
	fmt.Println(mySlice[0])
	mySlice[0] = 7
	fmt.Println(mySlice[0])
	mySlice[0]++ // up to 8
	fmt.Println(mySlice[0])
	mySlice[0]-- // back to 7
	fmt.Println(mySlice[0])
}
