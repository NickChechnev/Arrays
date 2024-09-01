package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func average(numbers ...float64) float64 {
	var sum float64
	for _, number := range numbers {
		sum += number
	}
	return sum / float64(len(numbers))
}

func main() {
	arguments := os.Args[1:]
	var nums []float64

	for _, argument := range arguments {
		number, err := strconv.ParseFloat(argument, 64)
		if err != nil {
			log.Fatal(err)
		}
		nums = append(nums, number)
	}
	fmt.Printf("Average is: %0.2f\n", average(nums...))
}
