package main

import (
	"fmt"
)

func main() {
	// init [key-type]value-type
	a := make(map[int]string)
	// map[key]val  syntax
	a[1] = "Trung"
	a[2] = "Phuc"
	fmt.Println(a)

	b := a[2]

	fmt.Println(b)

	// delete key and value

	delete(a, 1)
	fmt.Println(a)

	// check
	_, prs := a[2]
	fmt.Println(prs)

	c := map[string]int{"Trung": 10, "Phuc": 5, "Tung": 4}
	fmt.Println(c)
}
