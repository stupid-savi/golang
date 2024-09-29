package main

import "fmt"

func main() {

	// arrays are not used lot in Golang. Because to use arrays we need to assign it's length while declaring it
	// If length of array known use arrays else we use Slices which generates length dynamically
	// Benefits of using array :- Fixed Size , that is predictable, Memory optimization, faster access time

	var fruits [5]string
	// default values are known as zeroed values
	// Zeroed values for string-> empty string, int-> 0, bool-> false
	fmt.Println(fruits)
	fmt.Println(len(fruits))
	fruits[0] = "Orange"
	fmt.Println(fruits)
	fmt.Println(fruits[0])
	fmt.Println(fruits[0])
	// fruits[9] = "Apple" // give error since length is 5
	// assigning values while declaring

	arr := [3]bool{true, false} // zeroed value will add for 3rd element
	fmt.Println(arr)

	var arr2 = [3]int{1, 2} // Muttable values can't be declared with const in GO
	fmt.Println(arr2)

	// 2d array

	array2d := [3][4]int{{1, 2}, {5, 6}}
	fmt.Println(array2d)
	fmt.Println(len(array2d))
	fmt.Println(array2d[0])

	var t [2][3]int = [2][3]int{{}, {}}
	fmt.Println(t)

}
