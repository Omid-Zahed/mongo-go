//function argument
package main

import (
	"fmt"
)

func main() {
	greeting("omid", 22)
}
func greeting(name string, age int) {
	fmt.Printf("hi %v , your age is %v", name, age)

}
