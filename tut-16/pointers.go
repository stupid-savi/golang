package main

import "fmt"

func changeNum(num int) {
	num = 5
	println("changing inside function", num)
}

func changeNumWithPointer(num *int) {

	*num = 5 // to change the original value using *
	fmt.Println("The inner value...", *num)

}

func main() {
	// pointer are a reference to the memory address of a variable
	num := 1
	changeNum(num)
	fmt.Println("inside main function num", num)

	// the value of num isn't chnaged but why? when we pass variables as parameters there are passed as values not reference so it means inside changeNum function only value of num variable i.e 1 is passed and
	//inside the function a new variable num is being made 5 which was earlier storing the copy value 1
	// So to make outer num chnage we need to pass th memory address/reference of that variable not the value
	// pointer helps to pass the memory address

	// print the reference

	fmt.Println("the memory address:", &num, "the value can be print from memory addres using * operator:", *&num)

	changeNumWithPointer(&num)
	fmt.Println("inside main function num", num)

}
