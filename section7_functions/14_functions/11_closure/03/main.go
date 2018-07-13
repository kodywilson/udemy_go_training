package main

import "fmt"

func main() {
	x := 0
	increment := func() int {
		x++
		return x
	}
	for i := 10; i > 0; i-- {
		fmt.Println("Up!", increment(), "Down!", i)
	}
}
