package main

import "fmt"

var a int

type keg int
type cask string
type bottle bool

var b keg
var c cask
var d bottle

func main() {
	a = 42
	b = 43
	fmt.Println(a)
	fmt.Printf("%T\n", a)
	fmt.Println(b)
	fmt.Printf("%T\n", b)
	a = int(b)
	fmt.Println(a)
	fmt.Printf("%T\n", a)
	c = "Yo!"
	fmt.Println(c)
	d = true
	fmt.Println(d)
	c = "Ho Ho!"
	fmt.Println(c)
	c = "and a bottle of rum!"
	fmt.Println(c)
}
