package main

import "fmt"

/*
*
Operasi Perbandingan
- Operasi perbandingan adalah operasi untuk membandingkan dua buah data
- Operasi perbandingan adalah operasi yang menghasilkan nilai boolean (benar atau salah)
- Jika hasil operasinya adalah benar, maka nilainya adalah true
- Jika hasil operasinya adalah salah, maka nilainya adalah false

Operator Perbandingan			Keterangan
- >								Lebih Dari
- <								Kurang Dari
- >=							Lebih Dari Sama Dengan
- <=							Kurang Dari Sama Dengan
- ==							Sama Dengan
- !=							Tidak Sama Dengan
*/
func main() {
	var name1 = "Adrian"
	var name2 = "Adrian"

	var result bool = name1 == name2
	var result1 bool = name1 != name2

	fmt.Println(result)  // true
	fmt.Println(result1) // false
}
