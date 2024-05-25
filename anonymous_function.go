package main

import "fmt"

// Blacklist adalah tipe untuk fungsi yang menerima string dan mengembalikan boolean
type Blacklist func(string) bool

// registerUser adalah fungsi yang menerima nama dan fungsi blacklist sebagai parameter
func registerUser(name string, blacklist Blacklist) {
	// Mengecek apakah nama terdaftar dalam blacklist
	if blacklist(name) {
		// Jika nama ada di blacklist, cetak pesan "You are blocked"
		fmt.Println("You are blocked", name)
	} else {
		// Jika nama tidak ada di blacklist, cetak pesan "Welcome"
		fmt.Println("Welcome", name)
	}
}

func main() {
	// Mendefinisikan anonymous function untuk blacklist
	blacklist := func(name string) bool {
		// Fungsi ini mengembalikan true jika nama adalah "anjing"
		return name == "anjing"
	}
	// Mendaftarkan pengguna dengan nama "adrian" dan memeriksa menggunakan fungsi blacklist
	registerUser("adrian", blacklist)
	// Mendaftarkan pengguna dengan nama "anjing" dan memeriksa menggunakan fungsi blacklist
	registerUser("anjing", blacklist)

	// registerUser dengan nama "anjing" menggunakan anonymous function secara langsung
	registerUser("anjing", func(name string) bool {
		// Fungsi anonim yang sama, mengembalikan true jika nama adalah "anjing"
		return name == "anjing"
	}) // Output: You are blocked anjing

}
