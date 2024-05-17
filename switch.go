package main

import "fmt"

func main() {
	name := "Adrian"

	switch name {
	case "Adrian":
		fmt.Println("Hello Adrian")
	case "Naizar":
		fmt.Println("Hello Naizar")
	default:
		fmt.Println("Hi, Boleh Kenalan?")
	}

	// Switch deangan Short Statement
	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Nama terlalu panjang")
	case false:
		fmt.Println("Nama sudah benar")
	}

	// Switch tanpa Kondisi (namun lebih disarankan menggunakan if-statement karena lebih familiar)
	length := len(name)
	switch {
	case length > 10:
		fmt.Println("Nama terlalu panjang")
	case length > 5:
		fmt.Println("Nama lumayan panjang")
	default:
		fmt.Println("Nama sudah benar")

	}
}
