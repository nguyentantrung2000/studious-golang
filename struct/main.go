package main

import (
	"fmt"
)

type person struct {
	name string
	age  int
}

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 42

	return &p
}
func main() {
	fmt.Println(person{"Trung", 20})

	fmt.Println(person{name: "Bob", age: 50})

	fmt.Println(person{name: "Alice"})

	fmt.Println(&person{name: "Nhat", age: 40})

	fmt.Println(newPerson("Huu"))

	s := person{name: "Tung", age: 50}
	fmt.Println(s.age)

	sp := &s

	fmt.Println(sp)

	dog := struct {
		name   string
		isGood bool
	}{
		"Misa",
		true,
	}
	fmt.Println(dog)
}
