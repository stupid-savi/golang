package main

import "fmt"

func counter() func() int {
	count := 0
	return func() int {
		count += 1
		return count
	}
}

func main() {
	// closures is a concept by which if a functions return a function and the returned function is using any variable which is in the parent scope of it i.e is in outer function then the inner returned function keep the track of that variable even if the outer function is deleted from call stack. same as JS.

	count := counter()
	fmt.Println(count())
	fmt.Println(count())
	fmt.Println(count())
	fmt.Println(count())
	fmt.Println(count())
	fmt.Println(count())

}
