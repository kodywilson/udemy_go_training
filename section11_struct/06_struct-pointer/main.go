package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	p1 := &person{"James", 20}
	fmt.Println(p1)
	fmt.Printf("%T\n", p1)
	fmt.Println(p1.name)
	fmt.Println(p1.age)

	p2 := person{"Bond", 30}
	fmt.Println(p2)
	fmt.Printf("%T\n", p2)
	fmt.Println(p2.name)
	fmt.Println(p2.age)

	p3 := person{"Money", 24}
	fmt.Println(p3)
	fmt.Printf("%T\n", p3)
	fmt.Println(p3.name)
	fmt.Println(p3.age)

	p4 := person{"Penny", 34}
	fmt.Println(p4)
	fmt.Printf("%T\n", p4)
	fmt.Println(p4.name)
	fmt.Println(p4.age)
}
