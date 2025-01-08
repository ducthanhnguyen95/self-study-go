package main

import (
	"fmt"
	"github.com/ducthanhnguyen95/self-study-go/magazine"
)

func main() {
	//var s magazine.Subscriber
	//s.Rate = 4.99
	//fmt.Println(s.Rate)

	//subscriber := magazine.Subscriber{Name: "Aman Singh", Rate: 4.99, Active: true}
	//fmt.Println("Name:", subscriber.Name)
	//fmt.Println("Rate:", subscriber.Rate)
	//fmt.Println("Active:", subscriber.Active)

	subscriber := magazine.Subscriber{Rate: 4.99}
	fmt.Println("Name:", subscriber.Name)
	fmt.Println("Rate:", subscriber.Rate)
	fmt.Println("Active:", subscriber.Active)
}
