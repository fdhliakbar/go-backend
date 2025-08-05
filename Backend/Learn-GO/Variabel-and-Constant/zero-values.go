package main

import "fmt"

func main() {
	// Zero value adalah nilai default dari sebuah variabel, jika tidak diinisialisasi.
	// Setiap tipe data di Go memiliki zero value-nya masing-masing.
	
	var a int8
	var b int16
	var c int32
	var d int64
	var e float32
	var f float64
	var g bool
	var h string

	fmt.Println("Zero value dari a:", a)
	fmt.Println("Zero value dari b:", b)
	fmt.Println("Zero value dari c:", c)
	fmt.Println("Zero value dari d:", d)
	fmt.Println("Zero value dari e:", e)
	fmt.Println("Zero value dari f:", f)
	fmt.Println("Zero value dari g:", g)
	fmt.Println("Zero value dari h:", h)

}