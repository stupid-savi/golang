package main

import "fmt"

func main() {
	// string
	// var myName string = "Stupid Savi"
	// var myName = "Stupid Savi" // type infer
	myName := "Stupid Savi" // declaration and initialization/assignment at same time
	fmt.Println(myName)

	var userType string // declaration, type is mandatory here
	userType = "Admin"
	fmt.Println(userType)

	//int
	// var num int = 876
	// var num = 876
	num := 876
	fmt.Println(num)

	// float
	// var price float64 = 23.7898
	// var price = 23.7898
	price := 23.7898
	fmt.Println(price)

	//bool
	// var isPerson bool = true
	// var isPerson = true
	isPerson := true
	fmt.Println(isPerson)
}
