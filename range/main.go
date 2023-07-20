package main

import (
	"fmt"
)

func main() {

	// slices
	nums := []int{1, 2, 3}
	sum := 0

	// do not need index so blank
	for _, num := range nums {
		sum += num
	}
	fmt.Println(sum)

	// index
	for i, num := range nums {
		if num == 3 {

			// print the index of where the condition true
			fmt.Println(i)
		}
	}

	kvs := map[int]string{1: "Trung", 2: "Phuc", 3: "Linh"}
	for i, name := range kvs {
		// the index will be the key of map
		fmt.Println(i, name)
	}
	for i := range kvs {
		fmt.Println("key", i)
	}

	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
