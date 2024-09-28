package main

import "fmt"

// For looping there is only For loop in Go, No while loop and other types of for loop

func main() {

	// While loop from For loop

	count := 1

	for count <= 3 {
		fmt.Println("Count = ", count)
		count = count + 1
	}

	// Infinite loop

	// for {

	// 	fmt.Println(">>>>>>>>>>>>>>>>>>>This is an Infinite Loop>>>>>>>>>>>>>>>>>>")
	// }

	// Classic For loop

	for i := 0; i <= 10; i++ {

		if i == 2 {
			continue // skip current iteration
		}

		if i == 5 {
			break // end for loops
		}
		fmt.Println(i)
	}

	// 1.22 -> Range

	for i := range 10 {
		fmt.Println((i)) // print 0 to 9 excluding 10
	}

}
