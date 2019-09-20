package main

import "fmt"

func wrapper() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}

func main() {
	increment := wrapper()
	for i := 100; i > 0; i-- {
		fmt.Println("Up!", increment(), "Down!", i)
	}
}
