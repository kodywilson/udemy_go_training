package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func (p person) fullName() string {
	return p.first + p.last
}

func (p person) reverseName() string {
	return p.last + p.first
}

func (p person) sillyName() string {
	if p.first == "James" {
		return p.first + "Bongos"
	} else {
		return p.first + "McDoogal"
	}

}

func main() {
	p1 := person{"James", "Bond", 20}
	p2 := person{"Miss", "Moneypenny", 18}
	p3 := person{"Nuck", "Chorris", 50}
	fmt.Println(p1.first, p1.last, p1.age)
	fmt.Println(p1.fullName())
	fmt.Println(p1.reverseName())
	fmt.Println(p1.sillyName())
	fmt.Println(p2.first, p2.last, p2.age)
	fmt.Println(p2.fullName())
	fmt.Println(p2.sillyName())
	fmt.Println(p3.fullName())
	fmt.Println(p3.sillyName())
}
