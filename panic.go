package main

import "fmt"

/*
*
Panic
- Panic function adalah function yang bisa kita gunakan untuk menghentikan program
- Panic function biasanya dipanggil ketika terjadi panic pada saat program kita berjalan
- Saat panic function dipanggil, program akan terhenti, namun defer function tetap akan dieksekusi
*/

func endApp() {
	fmt.Println("End app")

	// Progam Recover yang Benar
	message := recover()
	fmt.Println("terjadi panic", message)
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("Ups Error")
	}

	// Cara salah menggunakan recover
	//message := recover()
	//fmt.Println("terjadi panic", message)
}

func main() {
	runApp(true)
	fmt.Println("Adrian azzaa")
}
