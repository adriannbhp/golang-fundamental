package main

import "fmt"

func main() {
	// Membuat map untuk menyimpan informasi kontak
	contacts := map[string]string{
		"Adrian":  "08123456789",
		"Naizar":  "08234567890",
		"Bimo":    "08345678901",
		"Joko":    "08456789012",
		"Rahmat":  "08567890123",
		"Sandi":   "08678901234",
		"Haris":   "08789012345",
		"Surya":   "08890123456",
		"Firdaus": "08901234567",
	}

	// Menampilkan daftar kontak
	fmt.Println("Daftar Kontak")
	for name, number := range contacts {
		fmt.Println("Nama: %s, Nomor: %s\n", name, number)
	}

	// Mencari nomor kontak berdasarkan nama
	searchName := "Dilla"
	if number, ok := contacts[searchName]; ok {
		fmt.Println("Nomor %s adalah: %s\n", searchName, number)
	} else {
		fmt.Println("Kontrak dengan nama %s tidak ditemukan\n", searchName)
	}
}
