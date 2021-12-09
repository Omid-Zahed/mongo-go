package main

import "fmt"

func main() {
	add_(4, 3, 32)

	numbers := []int{1, 2, 3}
	add_(numbers...)

}
func add_(numbers ...int) {
	fmt.Println(numbers)

}
