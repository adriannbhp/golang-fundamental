package main

import "fmt"

func main() {
	counter := 1
	// Kode program for
	for counter <= 10 {
		fmt.Println("Perulangan ke ", counter)
		counter++
	}

	fmt.Println("Selesai")

	// For dengan Statement
	for counter := 1; counter <= 10; counter++ {
		fmt.Println("Perulangan ke", counter)
	}

	/**
	For Range
	- For bisa digunakan untuk melakukan iterasi terhadap semua data collection
	- Data collection contohnya Array, Slice, dan Map
	*/

	// Menggunakan For Loop
	names := []string{"Adrian", "Bimo", "Hernawan", "Pratama"}
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	// Menggunakan For Range
	for index, name := range names {
		fmt.Println("Index", index, "=", name)
	}

	// Menggunakan For Range tanpa Index
	for _, name := range names { // ubah index menjadi _
		fmt.Println(name)
	}
}
