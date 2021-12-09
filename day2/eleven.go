//type casting
package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	x, y := 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	i, _ := strconv.ParseInt("2", 10, 8)
	fmt.Println(x, y, int8(x+y), z, strconv.FormatInt(22, 10), i)
}
