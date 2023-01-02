package main

import (
	"fmt"
	"testWithGo/math"
)

func main() {
	sum := math.Sum([]int{10, -2, 3})
	if sum != 11 {
		msg := fmt.Sprintf("FAIL: wanted 11 but got %d", sum)
		panic(msg)
	}
}
