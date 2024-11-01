package main

import (
	"fmt"
)

type stack[T any] struct {
	elments []T
}

// func printElements(val []string) {
// 	for _, item := range val {
// 		fmt.Println(item)
// 	}
// }

// make it generic
// Can also use any in plave of interface{}
//
//	func printElements[T interface{}](val []T) {
//		for _, item := range val {
//			fmt.Println(item)
//		}
//	}
//
//	func printElements[T comparable](val []T) {
//		for _, item := range val {
//			fmt.Println(item)
//		}
//	}
func printElements[T int | string, V bool | string](val []T, val2 V) {
	for _, item := range val {
		fmt.Println(item)
	}
	fmt.Println("V", val2)
}

func main() {

	myElements := []string{"Hi", "Bye"}
	myIntElements := []int{1, 2, 3, 4}
	// myBoolElements := []bool{true, false}
	printElements(myElements, false)
	printElements(myIntElements, "Hello")
	// printElements(myBoolElements)

	myStack := stack[int]{
		elments: []int{1, 2, 3, 4, 5},
	}

	yourStack := stack[[]bool]{
		elments: [][]bool{{true}, {false}},
	}

	fmt.Println(myStack.elments)
	fmt.Println(yourStack.elments[0][0])

}
