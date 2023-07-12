package main

import (
	"fmt"
)

func main() {
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(i)
	// }

	// i := 3
	// for {
	// 	i = i + 1
	// 	fmt.Println(i)
	// }

	for i := 0; i <= 5; i++ {
		if i%2 == 0 {
			fmt.Println("even", i)
			continue
		}
		fmt.Println("odd", i)
	}
}
