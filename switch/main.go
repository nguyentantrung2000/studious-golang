package main

import (
	"fmt"
	"time"
)

func main() {
	i := 7
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	default:
		fmt.Println("don't have that case")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Monday:
		fmt.Println("Cuoi tuan")
	default:
		fmt.Println("Trong tuan")
	}

	// whatIAm is a function
	whatIAm := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I am Bool")
		case int:
			fmt.Println("I am Inter")
		default:
			fmt.Println("IDK!!!!", t)
		}
	}
	// call func to use
	whatIAm(false)
	whatIAm(12)
	whatIAm("haha")
}
