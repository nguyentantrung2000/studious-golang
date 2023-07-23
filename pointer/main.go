package main

import (
	"fmt"
)

func intVal(ival int) {
	ival = 0
}

func zeroptr(zerop *int) {
	*zerop = 0
}

func main() {
	i := 1
	fmt.Println(i)

	intVal(i)
	fmt.Println(i)

	zeroptr(&i)
	fmt.Println(i)

	fmt.Println(&i)

}
