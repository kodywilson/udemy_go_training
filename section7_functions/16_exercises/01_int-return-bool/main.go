package main

import "fmt"

func divideBool(x int) (int, bool) {
	even := false
	if x%2 == 0 {
		even = true
	}
	return x / 2, even
}

func main() {
	fmt.Println(divideBool(100))
}
