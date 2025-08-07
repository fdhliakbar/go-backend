package main

import "fmt"

func main() {
	kata := "Hello"

	for index, char := range kata {
		fmt.Printf("Index %d: Karakter %c\n", index, char)
	}
}
