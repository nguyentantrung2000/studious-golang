package main

import (
	"fmt"
)

func minus(a int, b int) int {
	return a - b
}

// if multiple value have same type
func plus(a, b, c int) int {

	return a + b + c

}

func main() {
	res := minus(7, 5)
	fmt.Println(res)
	res = plus(1, 2, 3)
	fmt.Println(res)
}
