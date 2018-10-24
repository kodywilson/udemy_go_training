package main

import (
	"encoding/json"
	"os"
)

type Person struct {
	First       string
	Last        string
	Age         int
	notExported int
}

func main() {
	p1 := Person{"James", "Bond", 32, 007}
	json.NewEncoder(os.Stdout).Encode(p1)
}
