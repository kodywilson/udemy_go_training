package main

import "fmt"

func main() {
	evens := 0
	evensSum := 0
	odds := 0
	oddsSums := 0
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			evens = evens + 1
			evensSum = evensSum + i
		}
		if i%2 != 0 {
			odds = odds + 1
			oddsSums = oddsSums + i
		}
	}
	fmt.Println("There are", evens, "even numbers and they total", evensSum)
	fmt.Println("There are", odds, "odd numbers and they total", oddsSums)
}
