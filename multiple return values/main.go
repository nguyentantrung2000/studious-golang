package main

import (
	"fmt"
)

func vals() (int, int) {
	return 1, 2
}

func vals2() (string, string) {
	return "Trung", "Bo"
}

func vals3() (int, string) {
	return 1, "Phuc"

}

func main() {
	a, b := vals()
	c, d := vals2()

	fmt.Println(a)
	fmt.Println(b)

	fmt.Println(c)
	fmt.Println(d)

	// If you only want a subset of the returned values, use the blank identifier _.
	e, _ := vals3()
	fmt.Println(e)
}
