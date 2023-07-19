package main

import "fmt"

func main() {

	// if you don't add number in [] it will be init slices
	var a []string

	// add "b" element to slices or many elements
	a = append(a, "b", "c", "d")

	fmt.Println(a[2])

	// create new slices base on size old slices
	b := make([]string, len(a))
	copy(b, a)

	fmt.Println("copy", b)
	fmt.Println(len(b))

	fmt.Println(b[0:1])

	// >= <
	fmt.Println(b[1:2])

	// <
	fmt.Println(b[:2])

	fmt.Println(b[1:])

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println(twoD)
}
