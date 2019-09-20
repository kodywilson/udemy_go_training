package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

type DoubleZero struct {
	Person
	LicenseToKill bool
}

type GoldFinger struct {
	Person
}

func (p Person) Greeting() {
	fmt.Println("I'm just a regular person.")
}

func (dz DoubleZero) Greeting() {
	fmt.Println("Miss Moneypenny, so good to see you.")
}

func (gf GoldFinger) Greeting() {
	fmt.Println("No, Mr. Bond, I expect you to die!")
}

func main() {
	p1 := Person{
		Name: "Ian Flemming",
		Age:  44,
	}

	p2 := DoubleZero{
		Person: Person{
			Name: "James Bond",
			Age:  23,
		},
		LicenseToKill: true,
	}

	p3 := GoldFinger{
		Person: Person{
			Name: "Auric Goldfinger",
			Age: 45,
		},
	}

	p1.Greeting()
	p2.Greeting()
	p2.Person.Greeting()
	p3.Greeting()
}
