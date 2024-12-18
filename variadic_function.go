package main

import "fmt"

/**
Variadic Function
- Parameter yang berada di posisi terakhir, memiliki kemampuan dijadikan sebuah varargs
- Varargs artinya datanya bisa menerima lebih dari satu input, atau dianggap saja semacal Array
- Apa bedanya dengan parameter biasa dengan tipe data Array ?
	- Jika parameter tipe Array, kita wajib membuat array terlebih dahulu sebelum mengirimkan
	  function
	- Jika parameter menggunakan varargs, kita bisa langsung mengirim datanya, jika lebih dari satu
      cukup gunakan tanda koma
*/

// sumAll mengembalikan jumlah total dari sejumlah argumen yang diberikan
func sumAll(numbers ...int) int {
	total := 0 // Akumulasi penjumlahan dimulai dari 0

	// Menggunakan For Range tanpa Index (_)
	for _, number := range numbers {
		total += number
	}

	return total
}

func main() {
	// Menghitung total pengeluaran dengan argumen terpisah
	fmt.Println("Total Pengeluaran (3 kategori):", sumAll(10, 10, 10)) // Menggunakan Variable Argument

	// Menggunakan slice untuk menghitung total pengeluaran
	foodExpenses := []int{200, 300, 400, 150}
	fmt.Println("Total Pengeluaran (3 kategori)", sumAll(foodExpenses...))

	transportExpenses := []int{100, 200, 300}
	fmt.Println("Total Pengeluaran (Transportasi):", sumAll(transportExpenses...))

	entertainmentExpenses := []int{50, 75, 100}
	fmt.Println("Total Pengeluaran (Hiburan)", sumAll(entertainmentExpenses...))

	// Menggabungkan semua pengeluaran
	totalExpenses := append(foodExpenses, transportExpenses...)
	totalExpenses = append(totalExpenses, entertainmentExpenses...)
	fmt.Println("Total Pengeluaran (Semua Kategori):", sumAll(totalExpenses...))
}
