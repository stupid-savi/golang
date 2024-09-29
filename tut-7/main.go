package main

import "fmt"

func main() {
	const age int = 18
	// if age >= 18 {
	// 	fmt.Println("Person is an adult with age ", age)
	// } else {
	// 	fmt.Println("Person is not an adult with age ", age)
	// }

	// if age >= 18 {
	// 	fmt.Println("Person is an adult with age ", age)
	// } else if age < 18 && age >= 10 {
	// 	fmt.Println("Person is teenager with age ", age)
	// } else {
	// 	fmt.Println("Person is a child with age ", age)
	// }

	// direct vaiable assignment in the if condition but can be used inside conditional blocks not outside of it
	const isValidUser = true
	if bankBalance := 0; bankBalance <= 0 {
		fmt.Println("You are broke!", bankBalance)
	} else if bankBalance >= 10000 && bankBalance <= 99999 {
		fmt.Println("You can still survive", bankBalance)
	} else if bankBalance <= 9999 || !isValidUser {
		fmt.Println("You can still broke unauthorised user", bankBalance)
	}

	if isLoogedIn, role := true, "ADMIN"; isLoogedIn && role == "ADMIN" {
		fmt.Println("Congratulations! ", role, "You are logged in!")
	} else if isLoogedIn && role == "SUBADMIN" {
		fmt.Println("Congratulations! ", role, "You are logged in!")
	} else if isLoogedIn {
		fmt.Println("Congratulations! ", "User", "You are logged in!")
	} else {
		fmt.Println("401 Unauthorised Error")
	}
	// Golang doesn't have ternary operator use if else for it
}
