package main

import "fmt"

func main() {

	a := 43

	fmt.Println(a)  // 43
	fmt.Println(&a) // 0xc4200161a0

	var b *int = &a
	fmt.Println(b)  // 0xc4200161a0
	fmt.Println(*b) // 43

	*b = 42        // change the value at this address to 42
	fmt.Println(a) // 42
}
