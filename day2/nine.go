//short variable declarations :=
package main

import "fmt"

func main() {
	omidName, omidAge := Omid()
	fmt.Printf("you name is %v and %v years old \n", omidName, omidAge)
	favoriteColor, job := "black", "bugMaker"
	fmt.Println(favoriteColor, job)
}

func Omid() (name string, age int) {
	return "omid", 26
}
