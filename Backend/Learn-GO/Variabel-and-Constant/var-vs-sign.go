package main

import "fmt"

func main() {
	//  Di go ada dua cara untuk mendeklarasikan variabel
	// 1. Menggunakan var
	var name string = "Fadhli Akbar"
	fmt.Println("Nama: ", name)

	// 2. Menggunakan short declaration
	age := 50
	fmt.Println("Umur: ", age)

	// 3. Deklarasi variabel tanpa inisialisasi
	var a string
	var b int
	var c bool
	var d float64
	fmt.Printf("a: %q, b: %d, c: %t, d: %f\n", a, b, c, d)

	// Perbedaan Var dan Sign
	// var bisa digunakan inside dan outside dari function
	// sign hanya bisa digunakan di dalam function
}
