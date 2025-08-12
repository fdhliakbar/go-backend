package main

import "fmt"

func main() {
	var a int = 58
	var b int = 100
	var ptr *int

	ptr = &a
	fmt.Println("Value of a:", *ptr)

	ptr = &b
	fmt.Println("Value of b:", *ptr)
}
