package main

import (
	"fmt"
	"strconv"
)

func testFun(a int, b string) string {
	return strconv.Itoa(a) + b
}

// functions in go return multiple value

func multipleVal(a, b int) (int, int, int) {
	return a, b, 3
}

// functions in go are first citizen they can be stored in variable, passed as an argument to other functions and reuturn from other function

var test = func(fn func(a int) int) func(a int) int {
	return fn
}

func fn1(a int) int {
	return a
}
func main() {
	num := testFun(1, "savi")
	fmt.Println(num)

	val1, val2, val3 := multipleVal(1, 2)
	fmt.Println(val1, val2, val3)
	data := test(fn1)
	fmt.Println(data(90777))
}
