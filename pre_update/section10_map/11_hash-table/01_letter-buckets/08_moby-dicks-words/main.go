package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
)

func main() {
	// get the Moby Dick book
	res, err := http.Get("http://www.gutenberg.org/files/2701/old/moby10b.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(res.Body)
	defer res.Body.Close()
	// Set the split function for the scanning operation.
	scanner.Split(bufio.ScanWords)
	// Count the words.
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
