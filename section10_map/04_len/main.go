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
	myOther["Jack"] = "Arrrr Matey!!"
	fmt.Println(myOther)
	fmt.Println(len(myOther))
}
