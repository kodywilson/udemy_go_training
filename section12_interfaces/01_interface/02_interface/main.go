package main

import "fmt"

type Square struct {
	side float64
}

func (z Square) area() float64 {
	return z.side * z.side
}

type Shape interface {
	area() float64
}

func info(z Shape) {
	fmt.Println(z)
	fmt.Println(z.area())
}

func main() {
	s := Square{10}
	info(s)
	y := Square{100}
	info(y)
	x := Square{33}
	info(x)
	a := Square{50}
	info(a)
	b := Square{1000}
	info(b)
	c := Square{9}
	info(c)
	d := Square{-3}
	info(d)
	e := Square{-9}
	info(e)
}
