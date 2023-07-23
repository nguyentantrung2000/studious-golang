package main

import (
	"fmt"
)

// Anonymous functions are useful when you want to define a function inline without having to name it.

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	nextSeq := intSeq()
	fmt.Println(nextSeq())
	fmt.Println(nextSeq())
	fmt.Println(nextSeq())

	nextSeqs := intSeq()
	fmt.Println(nextSeqs())
}
