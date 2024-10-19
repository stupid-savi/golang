package main

import (
	"fmt"
	"runtime"
)

func main() {

	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(nums)

	// for i := 0; i < len(nums); i++ {
	// 	fmt.Println(nums[i])
	// }

	// for i := range len(nums) {
	// 	fmt.Println(nums[i])
	// }

	sum := 0

	for index, num := range nums {
		fmt.Println(num, index, nums[index], num == nums[index])
		sum = sum + num
	}

	// index, item := range Slice

	fmt.Println(sum)

	maps := map[string]int{"id": 1, "ac": 123}

	for key, value := range maps {
		fmt.Println(key, value)
	}

	for k := range maps {
		fmt.Println(k) // only keys
	}

	// key,value := range

	// string

	// starting byte, unicode value := range string

	// unicode point of rune
	// starting byte of rune

	// ascii values 0 tom 255 fits in 1 byte
	// unicode values can go above 255 so take can take more than 1 byte i.e 2 or 3 and so on
	// Here 😅 takes 4 byte to store so starting position 0, 1, 2, 3 takes by 😅 and t has 4 then

	// The unicode value of emoji is 128517 which in binary is 111110110000000101 = 19 bits
	// however theoretically 19bits/8 = can be fit in 3 bytes but Go uses a fixed size representaion for rune i.e rune (int32) is always 4 bytes (32 bits) to accommodate the maximum range of Unicode code points (which require up to 21 bits). That's why it is taking 4 byte

	for i, c := range "😅tupid Savi" {
		fmt.Println("Starting Byte", i, "Unicode Value", string(c))
	}

	fmt.Println("System architecture:", runtime.GOARCH)

}
