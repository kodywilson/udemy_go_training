package main

import "fmt"

// starting list of greetings
var myGreeting = map[int]string{
	0: "Good morning!",
	1: "Bonjour!",
	2: "Buenos dias!",
	3: "Bongiorno!",
}

func main() {

	fmt.Println(myGreeting)
	delete(myGreeting, 2)
	fmt.Println(myGreeting)
	myGreeting[4] = "Arrrrr Matey!!"
	fmt.Println(myGreeting)
	myGreeting[2] = "Buenos dias!" // let's put it back
	fmt.Println(myGreeting)        // notice how the order can change
	cleaner(1, 2, 3, 4)
	fmt.Println(myGreeting)
}

// pass list of keys to delete
func cleaner(listy ...int) {
	for _, x := range listy {
		fmt.Println(x)
		delete(myGreeting, x)
	}
}
