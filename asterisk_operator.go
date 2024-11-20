package main

import "fmt"

/*
  - Saat kita mengubah variable pointer, maka yang berubah hana variable tersebut.
  - Semua variable yang mengacu ke data yang sama tidak akan berubah
  - Jika kita ingin mengubah selruh variable yang mengacu ke data tersebut, kita bisa
    menggunakan operator *
*/

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
	address1 := Address{"Subang", "Jawa Barat", "Indonesia"}
	address2 := &address1

	address2.City = "Bandung"

	fmt.Println(address1) // address1 tidak berubah {Bandung Jawa Barat Indonesia}

	fmt.Println(address2) // address2 berubah &{Bandung Jawa Barat Indonesia}

	//address2 = &Address{"Jakarta", "DKI Jakarta", "Indonesia"}
	// Menggunakan Asterisk Operator *
	*address2 = Address{"Jakarta", "DKI Jakarta", "Indonesia"}

	fmt.Println(address1) // address1 tidak berubah {Bandung Jawa Barat Indonesia}
	fmt.Println(address2) // address2 berubah &{Jakarta Jawa Barat Indonesia}

	// Program Pass by Reference

	/**
	- Pointer adalah kemampuuan membuat reference ke lokasi data di memory yang sama, tanpa
	menduplikasi data yang sudah ada
	- Sederhananya, dengan kemampuan pointer, kita bisa membuat pass by reference
	*/

	//address3 := address1
	//address4 := &address1 // Pointer
	//
	//address3.City = "Subang"
	//
	//fmt.Println(address3)
	//fmt.Println(address4)

}
