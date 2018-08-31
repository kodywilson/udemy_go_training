package main

import "fmt"

func main() {

	myGreeting := map[int]string{
		0: "Good morning!",
		1: "Bonjour!",
		2: "Buenos dias!",
		3: "Bongiorno!",
		4: "Yo ho ho!",
	}

	for key, val := range myGreeting {
		fmt.Println(key, " - ", val)
	}
	var myNumbers = map[int]int{}
	y := 10
	for x := 0; x < 11; x++ {
		myNumbers[x] = y
		y--
	}
	for key, val := range myNumbers {
		if key%2 == 0 {
			fmt.Println(key, "is even and", val, "probably is too!")
		}
		if key%2 != 0 {
			fmt.Println(key, "is odd and", val, "probably is too!")
		}
	}
}
