package main

import "fmt"

func main() {

	var myGreeting = map[string]string{}
	myGreeting["Terry"] = "Hola!"
	myGreeting["Ernesto"] = "Good morning!"

	fmt.Println(myGreeting)

	var myOther = map[string]string{
		"Tim":  "Hola!",
		"Bert": "Good morning!",
	}
	fmt.Println(myOther)

	var myMixed = map[int]string{
		0: "Bill",
		2: "Tom",
		1: "Jill",
	}
	fmt.Println(myMixed)
}
