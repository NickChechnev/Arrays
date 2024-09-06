package main

import (
	"arrays/datafile"
	"fmt"
	"log"
)

func main() {
	lines, err := datafile.GetStrings("votes.txt")
	if err != nil {
		log.Fatal(err)
	}
	counts := make(map[string]int)
	for _, line := range lines {
		counts[line]++
	}
	for count, number := range counts {
		fmt.Printf("%s: %d\n", count, number)
	}
}
