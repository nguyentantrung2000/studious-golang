package main

import (
	"fmt"
)

const name = "Nguyen Tan Trung"

func main() {
	fmt.Println(name)

	const a = 10

	fmt.Println(10 + a)

	// A numeric constant has no type until it’s given one, such as by an explicit conversion.

	const d = 3e2 / a
	fmt.Println(int64(d))

}
