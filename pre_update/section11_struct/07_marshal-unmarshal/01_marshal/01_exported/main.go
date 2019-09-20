package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	First       string
	Last        string
	Age         int
	notExported int
}

func main() {
	p1 := Person{"James", "Bond", 32, 007}
	bs, _ := json.Marshal(p1)
	fmt.Println(bs)
	fmt.Printf("%T \n", bs)
	fmt.Println(string(bs))
	p2 := Person{"Money", "Penny", 24, 000}
	bs2, _ := json.Marshal(p2)
	fmt.Println(bs2)
	fmt.Printf("%T \n", bs2)
	fmt.Println(string(bs2))
}
