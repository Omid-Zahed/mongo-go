//multi result function
// user ( _ )
// diff = and :=

package main

import "fmt"

func main() {
	res, _ := omid(10, 20)
	fmt.Println(res)
}
func omid(x, y int) (int, string) {
	return x + y, "add"
}
