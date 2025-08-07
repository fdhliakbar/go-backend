package main

import "fmt"

func main() {

	for index := 1; index <= 10; index++ {
		if index == 5 {
			fmt.Println("Melanjutkan ke iterasi berikutnya")
			continue
		}
		fmt.Println("Iterasi ke-", index)
	}
}
