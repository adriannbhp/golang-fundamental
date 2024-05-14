package main

import "fmt"

func main() {
	var names [4]string

	names[0] = "Adrian"
	names[1] = "Bimo"
	names[2] = "Hernawan"
	names[3] = "Pratama"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])
	fmt.Println(names[3])

	// names[4] = "Ganteng" // ERROR invalid argument: index 4 out of bounds [0:4]

	/**
	Membuat Array Langsung
	- Di Go-Lang kita juga bisa membuat Array secara langsung saat deklarasi variable
	*/

	values := [...]int{ // [...] bisa menambahkan array tanpa batas tetapi harus langsung di deklarasikan
		90,
		81,
		79,
		60,
		56,
	}

	fmt.Println(values)
	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[2])

	// Di Go-Lang tidak dapat menghapus data di array tetapi hanya bisa mengosongkan array

	/**
	Function Array					Keterangan
	- len(array)					Untuk mendapatkan panjang Array
	- array[index]					Mendapatkan data di posisi index
	- array[index] = value			Mengubah data di posisi index
	*/
	fmt.Println(values)
	fmt.Println(len(values)) // Menghitung panjang array
	values[0] = 100          // Mengubah array di index 0 menjadi 100
	fmt.Println(values)
}
