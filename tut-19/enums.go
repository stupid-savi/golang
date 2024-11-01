package main

import "fmt"

// go doesn't have inbuilt enums we have to create them using type and const

type OrderStatus string

type MyRole int

const (
	InvoiceGenrated OrderStatus = "Invoice Genrated"
	Prepared                    = "Prepared"
	Delivered                   = "Delivered"
	Cancelled                   = "Cancelled"
	Returned                    = "Returned"
)

const (
	User MyRole = iota
	Admin
	Editor
	Author
	SuperAdmin
)

func changeOrderStatus(status OrderStatus) {
	fmt.Println("The order status is changed to", status)
}

func changeUserRole(role MyRole) {
	fmt.Println("ROle is changed to", role)
}

func main() {
	// control space for auto suggest
	changeOrderStatus(Returned)
	changeUserRole(Admin)
}
