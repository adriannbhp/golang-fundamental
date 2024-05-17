package main

import "fmt"

/*
*
Tipe Data Map
  - Pada Array atau Slice, untuk mengakses data, kita menggunakan index Number dimulai dari 0
  - Map adalah tipe data lain yang berisikan kumpulan data yang sama, namun kita bisa menentukan
    jenis tipe data index yang akan kita gunakan
  - Sederhananya, Map adalah tipe data kumpulan key-value (kata kunci - nilai), dimana kata kuncinya
    bersifat unik, tidak boleh sama
  - Berbeda dengan Array dan Slice, jumlah data yang kita masukkan ke dalam Map boleh sebanyak-banyaknya,
    asalkan kata kunci nya berbeda, jika kita gunakan kata kunci sama, maka secara otomatis data sebelumnya
    akan diganti dengan data baru
*/
func main() {

	//var person map[string] string = map[string]string{}
	//person["name"] = "Adrian"
	//person["address"] = "Bandung"

	person := map[string]string{
		"name":    "Adrian",
		"address": "Bandung",
	}

	fmt.Println(person["name"])
	fmt.Println(person["address"])
	fmt.Println(person)

	/**
	Function Map						Keterangan
	- len(map)							Untuk mendapatkan jumlah data di map
	- map[key]							Mengambil data di map dengan key
	- map[key] = value					MEngubah data di map dengan key
	- make(map[TypeKey]TypeValue)		Membuat map baru
	- delete(map, key)					Menghapus data di map dengan key
	*/

	// Membuat map baru untuk menyimpan informasi buku
	book := make(map[string]string)
	book["title"] = "Buku Golang"
	book["author"] = "Adrian"
	book["ups"] = "Salah"

	// Menampilkan isi map buku sebelum penghapusan
	fmt.Println("Buku sebelum dihapus:", book)

	// Menghapus kunci "ups" dari map buku
	delete(book, "ups")

	// Menampilkan isi map buku setelah penghapusan
	fmt.Println("Buku setelah dihapus:", book)
}
