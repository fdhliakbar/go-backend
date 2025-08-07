package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	var n uint8 = 0

	for n < 10 {
		fmt.Println(n)

		if n == 5 {
			fmt.Println("Breaking out of the loop at n =", n)
			break
		}
		n++
	}

	var n2 = 0

	for {
		fmt.Println(n2)
		n2++
		if n2 == 10 {
			fmt.Println("Break", n2)
			break
		} else if n2 >= 10 {
			break
		}
	}
}
