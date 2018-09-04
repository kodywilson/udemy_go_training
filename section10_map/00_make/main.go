package main

import "fmt"

func main() {

	var myGreeting = make(map[string]string)
	myGreeting["Terry"] = "Hola!"
	myGreeting["Ernesto"] = "Good morning!"

	fmt.Println(myGreeting)

	var myNumbers = make(map[int]int)
	for i := 0; i < 11; i++ {
		myNumbers[i] = i
	}
	fmt.Println(myNumbers)
}
