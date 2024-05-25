package main

import "fmt"

/**
Defer
- Defer function adalah function yang bisa kita jadwalkan untuk dieksekusi setelah sebuah function
  selesai di eksekusi
- Defer function akan selalu dieksekusi walaupun terjadi error di function yang dieksekusi
*/

func logging() {
	fmt.Println("Selesai memanggil function")
}

func runApplication() {
	defer logging()
	fmt.Println("Run application")
}
func main() {
	runApplication()
}
