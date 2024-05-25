package main

import "fmt"

func getHello(name string) string {
	return "Hello " + name
}

// Function untuk menghitung total harga belanjaan
func calculateTotalPrice(prices []int) int {
	total := 0
	for _, price := range prices {
		total += price
	}
	return total
}

func main() {
	result := getHello("Adrian")
	fmt.Println(result)

	// Tanpa menyimpan kedalam variabel
	fmt.Println(getHello("Naizar"))

	fmt.Println("-------------")
	fmt.Println("STUDI KASUS MENGHITUNG TOTAL HARGA BELANJAAN")

	// Slice Harga-harga barang belanjaan
	itemPrices := []int{10000, 15000, 20000, 25000, 30000}

	// Menampilkan harga-harga barang belanjaan menggunakan For Range
	fmt.Println("Harga Barang Belanjaan:")
	for i, price := range itemPrices {
		fmt.Printf("%d. Rp %d\n", i+1, price)
	}

	// Menghitung total harga belanjaan
	totalPrice := calculateTotalPrice(itemPrices)
	fmt.Printf("\nTotal Harga Belanjaan: Rp %d\n", totalPrice)

}
