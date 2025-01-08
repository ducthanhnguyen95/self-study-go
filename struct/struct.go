package main

import "fmt"

func main() {

	//var porsche car
	//porsche.name = "Porsche 911 R"
	//porsche.topSpeed = 323
	//fmt.Println("Name:", porsche.name)
	//fmt.Println("Top speed:", porsche.name)
	//
	//var bolts part
	//bolts.description = "Hex bolts"
	//bolts.count = 24
	//fmt.Println("Description:", bolts.description)
	//fmt.Println("Count:", bolts.count)

	//var bolts part
	//bolts.description = "Hex bolts"
	//bolts.count = 24
	//showInfo(bolts)

	//var subscriber1 subscriber
	//subscriber1.name = "Aman Singh"
	//fmt.Println("Name:", subscriber1.name)
	//var subscriber2 subscriber
	//subscriber2.name = "Beth Ryan"
	//fmt.Println("Name:", subscriber2.name)
	//p := minimumOrder("Hex bolts")
	//fmt.Println(p.description, p.count)

	subscriber1 := defaultSubscriber("Aman Singh")
	subscriber1.rate = 4.99
	printInfo(subscriber1)
	subscriber2 := defaultSubscriber("Beth Ryan")
	applyDiscount(&subscriber2)
	printInfo(subscriber2)

	//var value myStruct
	//value.myField = 3
	//var pointer *myStruct = &value
	//pointer.myField = 9
	//fmt.Println(pointer.myField)

}

type part struct {
	description string
	count       int
}
type car struct {
	name     string
	topSpeed float64
}

type subscriber struct {
	name   string
	rate   float64
	active bool
}

func showInfo(p part) {
	fmt.Println("Description:", p.description)
	fmt.Println("Count:", p.count)
}

func minimumOrder(description string) part {
	var p part
	p.description = description
	p.count = 100
	return p
}

func printInfo(s subscriber) {
	fmt.Println("Name:", s.name)
	fmt.Println("Monthly rate:", s.rate)
	fmt.Println("Active?", s.active)
}
func defaultSubscriber(name string) subscriber {
	var s subscriber
	s.name = name
	s.rate = 5.99
	s.active = true
	return s
}
func applyDiscount(s *subscriber) {
	s.rate = 4.99
}

type myStruct struct {
	myField int
}
