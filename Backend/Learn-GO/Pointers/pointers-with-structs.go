package main

import "fmt"

func main() {
	type Person struct {
		Name string
		Age  int
	}

	p1 := Person{Name: "Fadhli", Age: 30}
	p2 := Person{Name: "Akbar", Age: 25}

	ptr := &p1
	fmt.Println("Value of p1:", *ptr)

	ptr = &p2
	fmt.Println("Value of p2:", *ptr)
}
