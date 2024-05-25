package main

import "fmt"

/*
*
Struct
  - Struct adalah sebuah template data yang digunakan untuk mengabungkan nol atau lebih tipe daa
    lainnya dalam satu kesatuan
  - Struct biasanya representasi data dalam pogram aplikasi yang kita buat
  - Data di struck disimpan dalam field
  - Sederhananya struck adalah kumpulan dari field
*/

// Program Stuct
type Customer struct {
	Name, Address string
	Age           int
}

func main() {
	var adrian Customer
	fmt.Println(adrian)

	// Memanggil semua data di Struct
	adrian.Name = "Adrian Bimo"
	adrian.Address = "Indonesia"
	adrian.Age = 21
	fmt.Println(adrian)

	// Memanggil Struct 1 per satu
	fmt.Println(adrian.Name)
	fmt.Println(adrian.Address)
	fmt.Println(adrian.Age)

	// Struct Literals

	naizar := Customer{
		Name:    "Naizar",
		Address: "Indonesia",
		Age:     13,
	}
	fmt.Println(naizar)

	razka := Customer{Name: "Razka", Address: "Indonesia", Age: 10}
	fmt.Println(razka)
}
