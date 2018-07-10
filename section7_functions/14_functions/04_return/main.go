package main

import "fmt"

func main() {
	fmt.Println(greet("Kody ", "Wilson"))
}

func greet(fname, lname string) string {
	return fmt.Sprint(fname, lname)
}
