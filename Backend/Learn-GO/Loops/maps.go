package main

import "fmt"

func main() {
	person := map[string]string{
		"name": "Fadhli",
		"city": "Yogyakarta",
		"age":  "25",
		"job":  "Developer",
	}

	for key, value := range person {
		fmt.Printf("Key: %s, Value: %s\n", key, value)
	}
}
