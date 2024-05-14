package main

import "fmt"

func main() {
	/**
	Constant
	- Constant adalah variable yang nilainya tidak bisa diubah lagi setelah pertama kali diberi nilai
	- Cara pembuatan constant mirip dengan variable, yang membedakan hanya kata kunci yang
	  digunakan adalah const, bukan var
	- Saat pembuatan constant, kita wajib langsung menginisialisasikan datanya
	*/

	// Kode Program Constant
	const firstName string = "Adrian Bimo"
	const middleName = "Hernawan"
	const lastName = "Pratama"

	//firstName = "Adrian azza" // ERROR Cannot assign to firstName

	fmt.Println(firstName)

	/**
	Deklarasi Multipe Constant
	- Sama seperti variable, di Go-Lang juga kita bisa membuat constant secara sekaligus banyak
	*/

	const (
		firstName1  string = "Adrian Bimo" // Mendeklarasikan tipe data
		middleName1        = "Hernawan"    // Tanpa mendeklarasikan tipe data
		lastName1          = "Pratama"
	)

	fmt.Println(firstName1)
	fmt.Println(middleName1)
	fmt.Println(lastName1)
}
