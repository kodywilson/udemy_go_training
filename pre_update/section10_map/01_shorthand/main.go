package main

import "fmt"

func main() {

	myGreeting := make(map[string]string)
	myGreeting["Terry"] = "Hola!"
	myGreeting["Ernesto"] = "Good morning!"

	fmt.Println(myGreeting)

	myInts := make(map[int]int)
	myInts[0] = 0
	myInts[1] = 1
	myInts[2] = 2
	myInts[3] = 3
	myInts[myInts[3]*myInts[3]] = myInts[3] * myInts[3]
	fmt.Println(myInts)
	fmt.Println(myInts[1] + myInts[2])
	fmt.Println(myInts[3] * myInts[3])
	fmt.Println(myInts[9])
}
