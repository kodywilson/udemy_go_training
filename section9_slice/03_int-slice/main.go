package main

import "fmt"

func main() {
	a := make([]int, 0, 5)
	fmt.Println("---------------------")
	fmt.Println(a)
	fmt.Println(len(a))
	fmt.Println(cap(a))
	fmt.Println("---------------------")

	for i := 0; i < 80; i++ {
		a = append(a, i)
		fmt.Println("Len:", len(a), "Cap:", cap(a), "Val:", a[i])
	}
}
