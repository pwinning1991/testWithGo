package main

import (
	"fmt"
	"github.com/pwinning1991/testWithGo/math"
)

func main() {
	sum := math.Sum([]int{10, -2, 3})
	if sum != 11 {
		msg := fmt.Sprintf("FAIL: wanted 11 but got %d", sum)
		panic(msg)
	}
	fmt.Println("PASS")

	add := math.Add(5, 10)
	if add != 15 {
		msg := fmt.Sprintf("FAIL: wanted 15 but got %d", add)
		panic(msg)
	}
	fmt.Println("PASS")
}
