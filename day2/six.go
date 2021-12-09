//named return value
package main

import "fmt"

func main() {
	fmt.Println(test(20))
}
func test(number int) (y int, x int) {
	y = number
	x = 20
	return
}

