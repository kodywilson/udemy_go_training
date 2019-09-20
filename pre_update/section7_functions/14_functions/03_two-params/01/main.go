package main

import "fmt"

func main() {
	greet("Kody", "Wilson")
}

func greet(fname string, lname string) {
	fmt.Println(fname, lname)
}
