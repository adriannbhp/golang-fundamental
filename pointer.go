package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {

	/*
		Pass by Value
		  - Secara default di Go-Lang semua variable itu di passing, bukan by reference
		  - Artinya, jika kita mengirim sebuah variable ke dalam function, method atau variable lain,
		    sebenarnya yang dikirim adalah duplikasi value nya
	*/

	// Program Pass by Value
	address1 := Address{"Bandung", "Jawa Barat", "Indonesia"}
	address2 := address1

	address2.City = "Lembang"

	fmt.Println(address1) // address1 tidak berubah {Bandung Jawa Barat Indonesia}
	fmt.Println(address2) // address2 berubah {Lembang Jawa Barat Indonesia}

	// Program Pass by Reference

	/**
	- Pointer adalah kemampuuan membuat reference ke lokasi data di memory yang sama, tanpa
	menduplikasi data yang sudah ada
	- Sederhananya, dengan kemampuan pointer, kita bisa membuat pass by reference
	*/

	address3 := address1
	address4 := &address1 // Pointer

	address3.City = "Subang"

	fmt.Println(address3)
	fmt.Println(address4)

}
