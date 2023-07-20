package main

import (
	"fmt"
)

// take an arbitrary number of ints as arguments.
func sum(nums ...int) {
	fmt.Println(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func main() {
	sum(1, 2, 3)

	nums := []int{1, 2, 3}
	sum(nums...)
}
