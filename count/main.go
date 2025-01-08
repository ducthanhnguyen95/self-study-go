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
	fmt.Println(counts)
}

//func main() {
//	lines, err := datafile.GetStrings("D:/TCB/self-study-go/readfile/votes.txt")
//	if err != nil {
//		log.Fatal(err)
//	}
//	var names []string
//	var counts []int
//	for _, line := range lines {
//		matched := false
//		for i, name := range names {
//			if name == line {
//				counts[i]++
//				matched = true
//			}
//		}
//		if matched == false {
//			names = append(names, line)
//			counts = append(counts, 1)
//		}
//	}
//	for i, name := range names {
//		fmt.Printf("%s: %d\n", name, counts[i])
//	}
//}
