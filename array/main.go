package main

import "fmt"

func main() {
	// var notes [7]string
	// notes[0] = "do"
	// notes[1] = "re"
	// notes[2] = "mi"
	// fmt.Println(notes)
	// var dates [3]time.Time
	// dates[0] = time.Unix(1257894000, 0)
	// dates[1] = time.Unix(1447920000, 0)
	// dates[2] = time.Unix(1508632200, 0)
	// fmt.Println(dates[1])
	// var primes [5]int
	// primes[0] = 2
	// fmt.Println(primes[0])
	// fmt.Println(primes[2])
	// fmt.Println(primes[4])
	// fmt.Println(notes[3])
	// fmt.Println(notes[6])
	// fmt.Println(notes[0])

	// notes := [7]string{"do", "re", "mi", "fa", "so", "la", "ti"}
	// fmt.Println(notes[3], notes[6], notes[0])
	// primes := [5]int{2, 3, 5, 7, 11}
	// fmt.Println(primes[0], primes[2], primes[4])

	// text := [3]string{
	// 	"This is a series of long strings",
	// 	"which would be awkward to place",
	// 	"together on a single line",
	// }
	// fmt.Println(text)
	// fmt.Printf("%#v\n", text)
	// index := 1
	// fmt.Println(index, notes[index])
	// index = 3
	// fmt.Println(index, notes[index])
	// for i := 0; i <= 2; i++ {
	// 	fmt.Println(i, notes[i])
	// }
	// for i := 0; i <= 7; i++ {
	// 	fmt.Println(i, notes[i])
	// }

	notes := [7]string{"do", "re", "mi", "fa", "so", "la", "ti"}
	fmt.Println(len(notes))
	// for index, note := range notes {
	// 	fmt.Println(index, note)
	// }
	// for _, note := range notes {
	// 	fmt.Println(note)
	// }
	for index, _ := range notes {
		fmt.Println(index)
	}

}
