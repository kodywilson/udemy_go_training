package main

import "fmt"

func main() {

	myFriendsName := "Alf"

	switch {
	case len(myFriendsName) == 2:
		fmt.Println("Wassup my friend with name of length 2")
	case myFriendsName == "Bob":
		fmt.Println("Hi Bob")
	case myFriendsName == "Marcus":
		fmt.Println("Wassup Julian / Sushant")
	default:
		fmt.Println("nothing matched, this is the default")
	}
}
