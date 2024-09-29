package main

import (
	"fmt"
	"time"
)

func main() {
	const i int = 2
	switch i {
	case 1:
		fmt.Println("value is ", i)
	case 2:
		fmt.Println("value is ", i)
	case 3:
		fmt.Println("value is ", i)
	default:
		fmt.Println("other", i)
	}

	// default is optional, and no need to add break statement or return after case in go automatically considered

	// switch with multiple conditions

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's a weekned, Let's sleep")
	default:
		fmt.Println("Let's go to work")
	}

	// Type-switch
	// functions are first class citizen in Go. Which means you can pass functions as parameter to other functions and return them from other functions and save them in a varible also

	typeSwitchExample := func(m interface{}) {
		switch t := m.(type) {
		case string:
			fmt.Println("String", t)
		case int:
			fmt.Println("Intezer", t)
		case bool:
			fmt.Println("Boolean", t)
		default:
			fmt.Println("Other ", t)
		}

	}

	typeSwitchExample(1.2)

}
