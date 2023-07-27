package main

import (
	"fmt"
)

func main() {
	// make(chan val-type)
	messages := make(chan string)

	go func() {
		messages <- "ping"
	}()

	msg := <-messages

	fmt.Println(msg)
}
