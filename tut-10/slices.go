package main

import (
	"fmt"
	"slices"
)

func main() {
	// Slices are dynamic arrays
	// most used construct in Go

	// uninitialised slice is equal to nil (same as null in js )
	var nums []int
	fmt.Println(nums)
	fmt.Println(nums == nil)
	fmt.Println(len(nums))

	var nums2 = make([]int, 2)
	fmt.Println(nums2)
	fmt.Println(nums2 == nil)
	fmt.Println(cap(nums2))
	// if not specified the capacity equals to length
	// capacity -> maximum number of element can hold but in case of slice it's changes dynamically
	var nums3 = make([]int, 2, 5)
	nums3 = append(nums3, 1)
	fmt.Println(nums3)
	nums3 = append(nums3, 4)
	nums3 = append(nums3, 5)
	fmt.Println(nums3, cap(nums3))
	nums3 = append(nums3, 6)
	// capacity gets doubled i.e from 5 to 10
	fmt.Println(nums3, cap(nums3))

	nums4 := []int{1, 2, 3, 4, 5, 6}
	nums4 = append([]int{-2, -1, 0}, nums4...) // nums4... variadic operator = used for unpacking of slice
	fmt.Println(nums4)
	//If you want to add a new element to the slice, use the append function instead of directly assigning to an out-of-range index. Here's the corrected version of your code:
	// nums4[9] = 10 --- throw error index out of range
	// fmt.Println(nums4)

	// iusing index we can change item but only when length is specified

	var nums5 = make([]int, 2, 2)
	nums5[0] = 1
	nums5[1] = 2
	fmt.Println(nums5)

	// copying the slice

	var nums5Copy = make([]int, len(nums5))
	copy(nums5Copy, nums5)
	fmt.Println(nums5Copy)
	// slice operator
	numsSlice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	fmt.Println(numsSlice)
	fmt.Println(numsSlice[0:10]) // including 0 to 10-1 = 9 index
	fmt.Println(numsSlice[0:])   // including 0 to last element 0 to last
	fmt.Println(numsSlice[:10])  // from 0 to 10-1 = 9

	// if to value is present then exclude it

	var nums7 = []string{"savi", "riya", "bhavik"}
	var nums8 = []string{"savi", "riya", "bhavik"}
	fmt.Println(slices.Equal(nums7, nums8))
	// comparison will happen element by element and if at any level it becomes false then return false form there and ends the comparison

	twoDSlices := [][]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(twoDSlices)

}
