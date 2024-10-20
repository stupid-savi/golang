package main

import "fmt"

func sum(a ...int) int {

	sum := 0

	for _, num := range a {
		sum = sum + num
	}

	return sum

}

func variadicAnyType(num ...interface{}) {

	for _, item := range num {
		fmt.Println("item", item)

	}

}

func main() {

	// variadic functions are those functions which can take n number of arguments
	sum := sum(1, 2, 3, 4, 5, 6)
	fmt.Println(sum)
	variadicAnyType(1, 2, "savi", true)
	items := []interface{}{1, 2, 3, 4, "savi", true, 1.2, []int{0, 3, 4}, map[string]interface{}{"1": 1, "2": 2, "4": true}}
	// passing array for n number of items like spread in js
	variadicAnyType(items...)

}
