package main

import "fmt"

func main() {
	// numbers := make([]float64, 3)
	// numbers[0] = 19.7
	// numbers[2] = 25.2
	// for i, number := range numbers {
	// 	fmt.Println(i, number)
	// }
	// var letters = []string{"a", "b", "c"}
	// for i, letter := range letters {
	// 	fmt.Println(i, letter)
	// }
	//underlyingArray := [5]string{"a", "b", "c", "d", "e"}
	//slice4 := underlyingArray[:5]
	//fmt.Println(slice4)
	//array1 := [5]string{"a", "b", "c", "d", "e"}
	//slice1 := array1[0:3]
	//array1[1] = "X"
	//fmt.Println(array1)
	//fmt.Println(slice1)
	//slice := make([]int, 3, 5)
	//fmt.Println(slice)
	//slice := []string{"a", "b"}
	//var slice1 []string
	//fmt.Println(slice, len(slice))
	//slice = append(slice, "c")
	//fmt.Println(slice, len(slice))
	//slice1 = slice
	//slice = append(slice, "d", "e")
	//fmt.Println(slice, len(slice))
	//fmt.Println(slice1, len(slice1))
	//s1 := []string{"s1", "s1"}
	//s2 := append(s1, "s2", "s2")
	//s3 := append(s2, "s3", "s3")
	//s4 := append(s3, "s4", "s4")
	//fmt.Println("Length:", len(s1), "Capacity:", cap(s1))
	//fmt.Println("Length:", len(s2), "Capacity:", cap(s2))
	//fmt.Println("Length:", len(s3), "Capacity:", cap(s3))
	//fmt.Println("Length:", len(s4), "Capacity:", cap(s4))
	//fmt.Println(s1, s2, s3, s4)
	//s4[0] = "XX"
	//fmt.Println(s1, s2, s3, s4)
	var intSlice []int
	intSlice = append(intSlice)
	fmt.Printf("intSlice: %#v\n", intSlice)
	fmt.Println(intSlice)
}
