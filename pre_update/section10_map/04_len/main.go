package main

import "fmt"

func main() {

	var myGreeting = map[string]string{}
	myGreeting["Terry"] = "Hola!"
	myGreeting["Ernesto"] = "Good morning!"

	fmt.Println(myGreeting)
	fmt.Println("There are", len(myGreeting), "items in the map.")

	var myOther = map[string]string{
		"Tim":  "Hola!",
		"Bert": "Good morning!",
	}
	myOther["Jack"] = "Arrrr Matey!!"
	fmt.Println(myOther)
	fmt.Println("Now there are", len(myOther), "items in the map.")

	var myNumbers = map[int]int{}
	fmt.Println("myNumbers is", len(myNumbers), "long.")
	fmt.Println("It contains:", myNumbers)
	for i := 0; i < 100; i++ {
		myNumbers[i] = i
	}
	fmt.Println("myNumbers is now", len(myNumbers), "long.")
	fmt.Println("It contains:", myNumbers)
}
