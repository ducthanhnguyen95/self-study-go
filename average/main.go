package main

import (
	"fmt"
	"log"

	"github.com/ducthanhnguyen95/self-study-go/datafile"
)

func main() {
	numbers, err := datafile.GetFloats("D:/TCB/self-study-go/readfile/data.txt")
	if err != nil {
		log.Fatal(err)
	}
	var sum float64 = 0
	for _, number := range numbers {
		sum += number
	}
	sampleCount := float64(len(numbers))
	fmt.Printf("Average: %0.2f\n", sum/sampleCount)

}
