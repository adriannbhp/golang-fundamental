package main

import "fmt"

func main() {

	/**
	Type Declarations
	- Type Declarations adalah kemampuan membuat ulang tipe data baru dari tipe data yang sudah ada
	- Type Declarations biasanya digunakan untuk membuat alias terhadap tipe data yang sudah ada,
	  dengan tujuan agar lebih mudah dimengerti
	*/

	type NoKTP string // Deklarasi type

	var ktpAdrian NoKTP = "11111111" // Dengan Type Declaration tidak usah mendekalrasikan tipe data kembali

	var contoh string = "2222222"
	var contohKtp NoKTP = NoKTP(contoh) // Melakukan konversi

	fmt.Println(ktpAdrian)
	fmt.Println(contohKtp)

	// Studi Kasus Tanpa Type Declaration

	var suhuRuangan1 float64 = 25.5

	fmt.Println("Suhu Ruangan saat ini:", suhuRuangan1, "Derajat Celcius")

	// Studi Kasus Type Declartion

	type Temperatur float64

	var suhuRuangan Temperatur = 25.5

	fmt.Println("Suhu Ruangan saat ini:", suhuRuangan, "Derajat Celcius")
}
