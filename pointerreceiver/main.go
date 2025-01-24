package main

import "fmt"

type Number int

func (n *Number) Double() {
	*n *= 2
}
func main() {
	//number := Number(4)
	//fmt.Println("Original value of number:", number)
	//number.Double()
	//fmt.Println("number after calling Double:", number)

	//value := MyType(3)
	//pointer := &value
	//value.method()
	//value.pointerMethod()
	//pointer.method()
	//pointer.pointerMethod()

	soda := Liters(2)
	fmt.Printf("%0.3f liters equals %0.3f gallons\n", soda, soda.ToGallons())
	water := Milliliters(500)
	fmt.Printf("%0.3f milliliters equals %0.3f gallons\n", water, water.ToGallons())
}

type MyType int64

type Liters float64
type Milliliters float64
type Gallons float64

func (m MyType) method() {
	fmt.Println(m * 2)
}
func (m *MyType) pointerMethod() {
	fmt.Println(*m * 2)
}

func (l Liters) ToGallons() Gallons {
	return Gallons(l * 0.264)
}
func (m Milliliters) ToGallons() Gallons {
	return Gallons(m * 0.000264)
}
