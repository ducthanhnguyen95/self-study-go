// count tallies the number of times each line
// occurs within a file.
package main

import (
	"fmt"
	"github.com/ducthanhnguyen95/self-study-go/datafile"
	"log"
)

func main() {
	lines, err := datafile.GetStrings("D:/TCB/self-study-go/readfile/votes.txt")
	if err != nil {
		log.Fatal(err)
	}
	counts := make(map[string]int)
	for _, line := range lines {
		counts[line]++
	}
	//for name, count := range counts {
	//	fmt.Printf("Votes for %s: %d\n", name, count)
	//}
	fmt.Printf("%#v\n", myStruct)
	myStruct.number = 3.14
	myStruct.word = "pie"
	myStruct.toggle = true
	fmt.Println(myStruct.number)
	fmt.Println(myStruct.word)
	fmt.Println(myStruct.toggle)
	subscriber.name = "Aman Singh"
	subscriber.rate = 4.99
	subscriber.active = true
	fmt.Println("Name:", subscriber.name)
	fmt.Println("Monthly rate:", subscriber.rate)
	fmt.Println("Active?", subscriber.active)
}

var myStruct struct {
	number float64
	word   string
	toggle bool
}
var subscriber struct {
	name   string
	rate   float64
	active bool
}
