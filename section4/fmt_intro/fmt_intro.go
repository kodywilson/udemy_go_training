package main

import "fmt"

func main() {
	x := "This is only a test."
	fmt.Println("This is a test.")
	fmt.Println(x)
	fmt.Println("is of type")
	fmt.Printf("%T\n", x)
	fmt.Printf("The value of x is : ")
	fmt.Printf("%v\n", x)
	fmt.Println("Tab\tTab\tTab")

	y := fmt.Sprintf("%v\t%v\t%v\t", x, x, x)
	fmt.Println(y)
}
