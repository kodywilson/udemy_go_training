package main

import "fmt"

var x = 3

//z := 5 // Can not do that.
var z = 5 // This is how you do it.
var w int

func main() {
	fmt.Println(x) // This works, x is in scope.
	y := 4
	fmt.Println(y) // Ok, y is in scope here.
	foo()
	fmt.Println(w) // See the default value.
}

func foo() {
	fmt.Println(x) // This will work too.
	//fmt.Println(y) // This would not work, y is not in scope here.
	fmt.Println(z) //
	fmt.Println(x + z)
}
