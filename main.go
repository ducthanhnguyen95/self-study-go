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

	//subscriber := magazine.Subscriber{Rate: 4.99}
	//fmt.Println("Name:", subscriber.Name)
	//fmt.Println("Rate:", subscriber.Rate)
	//fmt.Println("Active:", subscriber.Active)

	//var employee magazine.Employee
	//employee.Name = "Jane"
	//employee.Salary = 60000
	//fmt.Println(employee.Name)
	//fmt.Println(employee.Salary)

	//var address magazine.Address
	//address.Street = "123 Oak St"
	//address.City = "Omaha"
	//address.State = "NE"
	//address.PostalCode = "68111"
	//fmt.Println(address)

	//address := magazine.Address{Street: "123 Oak St",
	//	City: "Omaha", State: "NE", PostalCode: "68111"}
	//subscriber := magazine.Subscriber{Name: "Aman Singh"}
	//subscriber.HomeAddress = address
	//fmt.Println(subscriber.HomeAddress)

	//subscriber := magazine.Subscriber{}
	//subscriber.HomeAddress.PostalCode = "68111"
	//fmt.Printf("%#v\n", subscriber.HomeAddress.PostalCode)

	//subscriber := magazine.Subscriber{Name: "Aman Singh"}
	//subscriber.Address.Street = "123 Oak St"
	//subscriber.Address.City = "Omaha"
	//subscriber.Address.State = "NE"
	//subscriber.Address.PostalCode = "68111"
	//fmt.Println("Subscriber Name:", subscriber.Name)
	//fmt.Println("Street:", subscriber.Address.Street)
	//fmt.Println("City:", subscriber.Address.City)
	//fmt.Println("State:", subscriber.Address.State)
	//fmt.Println("Postal Code:", subscriber.Address.PostalCode)
	//employee := magazine.Employee{Name: "Joy Carr"}
	//employee.Address.Street = "456 Elm St"
	//employee.Address.City = "Portland"
	//employee.Address.State = "OR"
	//employee.Address.PostalCode = "97222"
	//fmt.Println("Employee Name:", employee.Name)
	//fmt.Println("Street:", employee.Address.Street)
	//fmt.Println("City:", employee.Address.City)
	//fmt.Println("State:", employee.Address.State)
	//fmt.Println("Postal Code:", employee.Address.PostalCode)

	subscriber := magazine.Subscriber{Name: "Aman Singh"}
	subscriber.Street = "123 Oak St"
	subscriber.City = "Omaha"
	subscriber.State = "NE"
	subscriber.PostalCode = "68111"
	fmt.Println("Street:", subscriber.Street)
	fmt.Println("City:", subscriber.City)
	fmt.Println("State:", subscriber.State)
	fmt.Println("Postal Code:", subscriber.PostalCode)
	employee := magazine.Employee{Name: "Joy Carr"}
	employee.Street = "456 Elm St"
	employee.City = "Portland"
	employee.State = "OR"
	employee.PostalCode = "97222"
	fmt.Println("Street:", employee.Street)
	fmt.Println("City:", employee.City)
	fmt.Println("State:", employee.State)
	fmt.Println("Postal Code:", employee.PostalCode)
}
