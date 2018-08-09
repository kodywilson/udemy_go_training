package main

import "fmt"

func main() {
	a := []int{1, 3, 5, 7, 9, 11}
	fmt.Printf("%T\n", a)
	fmt.Println(len(a))
	fmt.Println(cap(a))
	fmt.Println(a)

	a = append(a, 13)
	fmt.Println(len(a))
	fmt.Println(cap(a))
	fmt.Println(a)
}
