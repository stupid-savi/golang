package main

import "fmt"

const age = 40

// name:= "rohane" Cab't use this syntax to declare variable outside of function

func main() {
	const person string = "savi"

	name := "rohan"

	//age = 25 constants can't be reassigned
	fmt.Println(person, age, name)

	// constant grouping

	const (
		lang   = "JS"
		amount = "500INR"
	)

	fmt.Println(lang, amount)

}
