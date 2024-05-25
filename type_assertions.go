package main

import "fmt"

/*
*
Type Assertions
- Type Assertions merupakan kemampuan merubah tipe data menjadi tipe data yang diingingkan
- Fitur ini sering sekali digunakan ketika kita bertemu dengan data interface kosong
*/

func random() interface{} {
	return 100
}

func main() {
	var result any = random()
	//var resultString string = result.(string)
	//fmt.Println(resultString)
	//
	//var resultInt int = result.(int)
	//fmt.Println(resultInt)

	/**
	Output:
	OK
	panic: interface conversion: interface {} is string, not int
	*/

	/* Type Assertions Menggunakan Switch
	   - Saat salah menggunakan type assertions, maka bisa berakibat terjadi panic di aplikasi kita
	   - Jika panic dan tidak ter-cover, maka otomatis program kita akan mati
	   - Agar lebih aman, sebaiknya kita menggunakan switch expression untuk melakukan type assertions
	*/

	switch value := result.(type) {
	case string:
		fmt.Println("String", value)
	case int:
		fmt.Println("Int", value)
	default:
		fmt.Println("Unknown", value)
	}
}
