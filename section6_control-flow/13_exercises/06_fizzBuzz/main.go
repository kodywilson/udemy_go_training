// Write a program that prints the numbers from 1 to 100.
// Multiples of three print "Fizz" instead of the number.
// Multiples of five print "Buzz".
// Multiples of both three and five print "FizzBuzz".
// I put the number too because I wanted to see if with the Fizz Buzz.

package main

import "fmt"

func main() {
	for i := 1; i < 101; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println(i, "FizzBuzz")
			continue
		}
		if i%3 == 0 {
			fmt.Println(i, "Fizz")
		}
		if i%5 == 0 {
			fmt.Println(i, "Buzz")
		}
	}
}
