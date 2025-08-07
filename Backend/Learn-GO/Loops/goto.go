package main

import "fmt"

func main() {
	i := 1

	for {
		if i == 3 {
			i++
			goto skip
		}
		if i > 5 {
			break
		}

		fmt.Println("i =", i)
		i++
		continue

	skip:
		fmt.Println("Lewatkan i = 3")
	}
}
