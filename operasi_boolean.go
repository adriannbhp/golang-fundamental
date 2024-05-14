package main

import "fmt"

/*
*
Operasi Boolean				Keterangan
- &&						Dan
- ||						Atau
- !							Kebalikan
*/

func main() {
	var nilaiAkhir = 90 // Menggunakan Var
	nilaiAbsensi := 80

	lulusNilaiAkhir := nilaiAkhir > 80
	lulusAbsensi := nilaiAbsensi > 80

	lulus := lulusNilaiAkhir && lulusAbsensi

	fmt.Println(lulus)
}
