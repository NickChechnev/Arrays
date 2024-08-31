package main

import "fmt"

func main() {
	numbers := []float64{71.8, 56.2, 89.5}
	var sum float64
	for _, number := range numbers {
		sum += number
	}
	fmt.Println(sum)
	fmt.Println("Average amount of meat is:", sum/float64(len(numbers)))
}
