package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5}

	for i := range numbers {
		fmt.Print(numbers[i])
	}
}
