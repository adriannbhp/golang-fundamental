package main

import "fmt"

/**
Variable
- Variable adalah tempat untuk menyimpan data
- Variable digunakan agar kita bisa mengakses data yang sama dimanapun kita mau
- Di Go-Lang Variable hanya bisa menyimpan tipe data yang sama, jika kita ingin
  menyimpan data yang berbeda-beda jenis, kita harus membuat beberapa variable
- Untuk membuat variable, kita bisa menggunakan kata kunci var, lalu diikuti dengan
  nama variable dan tipe datanya
*/

func main() {
	var name string

	name = "Adrian Bimo"
	fmt.Println(name)

	name = "Adrian Bimo Hernawan Pratama"
	fmt.Println(name)

	/**
	Tipe Data Variable
	- Saat kita membuat variable, maka kita wajib menyebutkan tipe data variable tersebut
	- Namun jika kita langsung menginisialisasikan data pada variable nya, maka kita tidak
	  wajib menyebutkan tipe data variable nya
	*/

	// Tanpa menyebutkan tipe data
	var name1 = "Adrianbhp"
	var age = 21
	fmt.Println(name1)
	fmt.Println(age)

	name1 = "Adrian ganteng"
	fmt.Println(name1)

	/**
	Kata Kunci Var
	- Di Go-Lang, kata kunci var saat membuat variable tidak lah wajib
	- Asalkan saat membuat variable kita langsung menginisialisasi datanya
	- Agar tidak perlu menggunakan kata kunci var, kita perlu menggunakan kata kunci := saat
	  menginisialisasikan data pada variable tersebut
	*/

	// Tanpa kata kunci var
	name2 := "Adrianbhp"
	fmt.Println(name2)

	name2 = "Adrian ganteng"
	fmt.Println(name2)

	/**
	Deklarasi Multipe Variable
	- Di Go-Lang kita bisa membuat variable secara sekaligus banyak
	- Code yang dibuat akan lebih bagus dan mudah dibaca
	*/

	var (
		firstName  = "Adrian Bimo"
		middleName = "Hernawan"
		lastName   = "Pratama"
	)

	fmt.Println(firstName)
	fmt.Println(middleName)
	fmt.Println(lastName)
}
