package main

import "fmt"

/**
Tipe Data Slice
- Tipe data Slice adalah potongan dari data Array
- Slice mirip dengan Array, yang membedakan adalah ukutan Slice bisa berubah
- Slice dan Array selalu terkoneksi, dimana Slice adalah data yang mengakses sebagian
  atau seluruh data di Array

Detail Tipe Data Slice
- Tipe Data Slice memiliki 3 data, yaitu pointer, length, dan capacity
- Pointer adalah penunjuk data pertama di array para slice
- Length adalah panjang dari slice, dan
- Capacity adalah kapasitas dari slice, dimana length tidak boleh lebih dari capacity

Membuat Slice Dari Array			Keterangan
- array[low:high]					Membuat slice dari array dimulai index low
									sampai index sebelum high

- array[low:]						Membuat slice dari array dimulai index low
									sampai index akhir di array

- array[:high]						Membuat slice dari array dimulai index 0
									sampai index sebelum high

- array[:]							Membuat slice dari array dimulai index 0
									sampai index akhir di array
*/

func main() {
	names := [...]string{"Adrian", "Bimo", "Hernawan", "Pratama", "Naizar", "Putra", "Herlino"} // Array

	slice1 := names[4:6]
	fmt.Println(slice1)

	slice2 := names[:4]
	fmt.Println(slice2)

	slice3 := names[4:]
	fmt.Println(slice3)

	slice4 := names[:]
	fmt.Println(slice4)

	/**
	Function Slice
	- len(slice)								Untuk mendapatkan panjang

	- cap(slice)								Untuk mendapat kapasitas

	- append(slice, data)						Membuat slice baru dengan menambah data ke posisi
												terakhir slice, jika kapasitas penuh, maka akan
												membuat array baru

	- make([]TypeData, lenght, capacity)		Membuat slice baru

	- copy(destination, source)					Menyalin slice dari source ke destination
	*/

	// Program Append Slice
	days := [...]string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu"} // Array

	daySlice1 := days[5:]  // Sabtu, Minggu
	fmt.Println(daySlice1) // [Sabtu Minggu]

	daySlice1[0] = "Sabtu Baru"
	daySlice1[1] = "Minggu Baru"
	fmt.Println(daySlice1) // [Sabtu Baru Minggu Baru]
	fmt.Println(days)      // [Senin Selasa Rabu Kamis Jumat Sabtu Baru Minggu Baru]

	daySlice2 := append(daySlice1, "Libur Baru")
	daySlice2[0] = "Sabtu Lama"
	// daysBaru := [...]string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu Lama", "Minggu Baru", "Libur Baru"}

	fmt.Println(daySlice1) // [Sabtu Baru Minggu Baru]
	fmt.Println(daySlice2) // [Sabtu Lama Minggu Baru Libur Baru]
	fmt.Println(days)      // [Senin Selasa Rabu Kamis Jumat Sabtu Baru Minggu Baru]

	// Kode Program Make Slice

	// Membuat slice dengan panjang 2 dan kapasitas 5
	newSlice := make([]string, 2, 5) // Slice

	// Mengisi elemen-elemen slice
	newSlice[0] = "Adrian"
	newSlice[1] = "Naizar"
	// newSlice[2] = "Adrian" // ERROR, harusnya menggunakan append

	// Menampilkan slice awal
	fmt.Println("Slice awal: ", newSlice)
	fmt.Println("Panjang awal: ", len(newSlice))
	fmt.Println("Kapasitas awal: ", cap(newSlice))

	// Menambahkan elemen baru menggunakan append
	newSlice = append(newSlice, "Dilla")
	newSlice = append(newSlice, "Amoy")

	// Menampilkan slice setelah penambahan elemen
	fmt.Println("Slice setelah append:", newSlice)
	fmt.Println("Panjang setelah append:", len(newSlice))
	fmt.Println("Kapasitas setelah append:", cap(newSlice))

	// Menambahkan lebih banyak elemen untuk melihat perubahan kapasitas
	newSlice = append(newSlice, "Safa")
	newSlice = append(newSlice, "Fiona")

	// Menampilkan slice setelah menambahkan lebih banya elemen
	fmt.Println("Slice setelah menambahkan lebih banyak elemen:", newSlice)
	fmt.Println("Panjang akhir:", len(newSlice))
	fmt.Println("Kapasitas akhir:", cap(newSlice))

	newSlice2 := append(newSlice, "Mika")
	fmt.Println(newSlice)
	fmt.Println(newSlice2)
	fmt.Println(len(newSlice2))
	fmt.Println(cap(newSlice2))

	newSlice2[0] = "Bimo" // Mengubah
	fmt.Println(newSlice2)
	fmt.Println(newSlice)

	// Program Copy Slice
	fromSlice := days[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice)) // Slice

	copy(toSlice, fromSlice)

	fmt.Println(fromSlice)
	fmt.Println(toSlice)

	fromNewSlice := newSlice[:]
	toSlice1 := make([]string, len(fromNewSlice), cap(fromNewSlice)) // Slice

	copy(toSlice1, fromNewSlice)

	fmt.Println(toSlice1)
	fmt.Println(fromNewSlice)

	fromNewSlice2 := newSlice2[:]
	toSlice2 := make([]string, len(fromNewSlice2), cap(fromNewSlice2)) // Slice

	copy(toSlice2, fromNewSlice2)

	fmt.Println(toSlice2)
	fmt.Println(fromNewSlice2)

	/**
	Hati-Hati Saat Membuat Array
	- Saat membuat Array, kita harus berhati-hati, jika salah, maka yang kita buat bukanlah Array,
	  melainkan Slice
	*/

	// Array vs Slice
	iniArray := [...]int{1, 2, 3, 4, 5}
	iniArray2 := [5]int{1, 2, 3, 4, 5}
	iniSlice := []int{1, 2, 3, 4, 5}

	fmt.Println(iniArray)
	fmt.Println(iniArray2)
	fmt.Println(iniSlice)

	// Di golang akan lebih sering menggunakan slice dibandingkan array

}
