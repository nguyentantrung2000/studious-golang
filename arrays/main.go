package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello Bo")

	// init array
	var a [5]int

	fmt.Println(a)

	var b [3]string

	fmt.Println(b)

	// array with elements

	c := [3]int{1, 2, 3}
	fmt.Println(c)

	d := [2]string{"a", "b"}
	fmt.Println(d)

	// two demension

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}

	}
	fmt.Println(twoD)

}
