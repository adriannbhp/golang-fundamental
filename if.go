package main

import "fmt"

func main() {
	name := "Bimo"

	if name == "Adrian" {
		fmt.Print("Hello Adrian")
	} else if name == "Naizar" {
		fmt.Println("Hello Naizar")
	} else {
		fmt.Println("Hai, boleh kenalan?")
	}

	// If dengan Short Statement
	if length := len(name); length > 5 {
		fmt.Println("Nama terlalu panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}
}
