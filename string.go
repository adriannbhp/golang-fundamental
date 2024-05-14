package main

import "fmt"

/**
Tipe Data String
- String adalah tipe data kumpulan karakter
- Jumlah karakter di dalam String bisa nol sampai tidak terhingga
- Tipe data String di Go-Lang direpresentasikan dengan kata kunci string
- Nilai data String di Go-Lang selal diawali dengan karakter" (Petika dua) dan
  diakhiri dengan karakter" (petik dua)
*/

/**
Function untuk String
Function						Keterangan
- len("string")					Menghitung jumlah karakter di String
- "string"[number]				Mengambil karakter pada posisi yang ditentukan
*/

func main() {
	fmt.Println("Adrian")
	fmt.Println("Bimo")
	fmt.Println("Hernawan")
	fmt.Println("Pratama")

	fmt.Println(len("Adrian Bimo"))
	fmt.Println("Adrian Bimo Hernawan"[0])
	fmt.Println("Adrian Bimo Hernawan Pratama"[1])

}
