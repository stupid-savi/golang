package main

import (
	"fmt"
	"time"
)

// structs are like structure with the help we can create any custom data type object or class

type order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time // nanoseconds precision
}

// convetion is use first letter of struct for creating a method init
func (o order) getAmount() float32 {
	return o.amount
}

func (o order) setAmount(amount float32) float32 {
	o.amount = amount // The amount won't get set because o is still a distinct copy of value not a same reference so we need to use pointers for that
	return o.amount
}

// using pointer
func (o *order) setAmount2(amount float32) {
	o.amount = amount // we doesnot need to use *o.amount = amount just like in the normal function example because the struct itself do dereference for us
}

// Note:-  use pointer when we need to mutate data not to get data

// We can't make field required in the struct directly but can add this check in constructor function

// Constructor

func NewOrder(id string, amount float32, createdAt time.Time) *order {

	myorder := order{
		id:        id,
		amount:    amount,
		createdAt: createdAt,
	}

	return &myorder

}

func (o *order) changeStatus(status string) {
	o.status = status
}

func main() {

	myorder := order{
		id:        "1",
		amount:    100.00,
		status:    "Shipped",
		createdAt: time.Now(),
	}

	myOrder2 := order{
		id:        "2",
		amount:    109.98,
		createdAt: time.Now(),
	}

	// Since there is no mandatory fields in the struct and if we don't provide them then they will we replaced with zeroed values. int 0 float 0 bool false string ""

	myOrder2.status = "Delivered"
	myorder.status = "Returned"
	// can chage properties here just like a class initialised object

	fmt.Println("myorder", myorder)
	fmt.Println("myorder", myOrder2)
	myOrder2.setAmount(180.87) // amount didn't get set
	myOrder2.setAmount2(500.89)

	fmt.Println(myOrder2.getAmount(), myOrder2)

	myOrder3 := NewOrder("3", 789.09, time.Now())

	fmt.Println(myOrder3)
	myOrder3.status = "Shipped"
	fmt.Println(myOrder3)

	// direct struct initiation

	myStruct := struct {
		name string
		role string
	}{"savi", "Software Engineer"}
	fmt.Println(myStruct.name, myStruct.role, myStruct)

}
