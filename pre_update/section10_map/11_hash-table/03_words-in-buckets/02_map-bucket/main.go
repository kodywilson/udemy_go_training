package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
)

func main() {
	// get the Moby Dick book
	res, err := http.Get("http://www.gutenberg.org/cache/epub/1661/pg1661.txt")
	if err != nil {
		log.Fatal(err)
	}
	// scan the page
	scanner := bufio.NewScanner(res.Body)
	defer res.Body.Close()
	// Set the split function for the scanning operation.
	scanner.Split(bufio.ScanWords)
	// Create map with a key of int and a value of another map with a key of
	// string which will be the word and a value of int, which will be number
	// of times that word occurs
	buckets := make(map[int]map[string]int)
	// create slices to hold words
	for i := 0; i < 24; i++ {
		buckets[i] = make(map[string]int)
	}
	// Loop over the words
	for scanner.Scan() {
		word := scanner.Text()
		n := HashBucket(word, 24)
		buckets[n][word]++
	}
	// Print the words in one of the buckets
	for k, v := range buckets[6] {
		fmt.Println(v, " \t- ", k)
	}
}

func HashBucket(word string, buckets int) int {
	var sum int
	for _, v := range word {
		sum += int(v)
	}
	return sum % buckets
}
