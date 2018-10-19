package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	First string
	Last  string
	Age   int
}

func main() {
	var p1 Person
	fmt.Println(p1.First)
	fmt.Println(p1.Last)
	fmt.Println(p1.Age)

	bs := []byte(`{"First":"James", "Last":"Bond", "Age":32}`)
	json.Unmarshal(bs, &p1)

	fmt.Println("----------------")
	fmt.Println(p1.First)
	fmt.Println(p1.Last)
	fmt.Println(p1.Age)
	fmt.Printf("%T \n", p1)

	bs2 := []byte(`{"First":"Money", "Last":"Penny", "Age":24}`)
	json.Unmarshal(bs2, &p1)

	fmt.Println("----------------")
	fmt.Println(p1.First)
	fmt.Println(p1.Last)
	fmt.Println(p1.Age)
	fmt.Printf("%T \n", p1)

	bs3 := []byte(`{"First":"Nuck", "Last":"Chorris", "Age":55}`)
	json.Unmarshal(bs3, &p1)

	fmt.Println("----------------")
	fmt.Println(p1.First)
	fmt.Println(p1.Last)
	fmt.Println(p1.Age)
	fmt.Printf("%T \n", p1)

	bs4 := []byte(`{"First":"Silly", "Last":"String", "Age":33}`)
	json.Unmarshal(bs4, &p1)

	fmt.Println("----------------")
	fmt.Println(p1.First)
	fmt.Println(p1.Last)
	fmt.Println(p1.Age)
	fmt.Printf("%T \n", p1)
}
