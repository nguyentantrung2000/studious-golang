package main

import (
	"fmt"
)

func main() {
	const a = 10
	if a >= 10 {
		fmt.Println("a greater than or equal value")
	} else {
		fmt.Println("smaller than")
	}

	const num = 899
	if num < 0 {
		fmt.Println("negative")
	} else if num >= 100 {
		fmt.Println("have multiple digits")
	} else {
		fmt.Println("2 digits")
	}
}
