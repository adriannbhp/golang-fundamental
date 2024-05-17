package main

import "fmt"

// getCompleteName() mengembalikan tiga bagian nama lengkap karyawan.
func getCompleteName() (firstName, middleName, lastName string) {
	firstName = "Adrian Bimo"
	middleName = "Hernawan"
	lastName = "Pratama"
	return firstName, middleName, lastName
}

func main() {
	// Memanggil fungsi getCompleteName dan menyimpan hasil pengembaliannya
	firstName, middleName, lastName := getCompleteName()

	// Menampilkan nama lengkap
	fmt.Printf("Nama Lengkap: %s %s %s\n", firstName, middleName, lastName)
}
