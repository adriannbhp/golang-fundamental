package main

import "fmt"

func getGoodBye(name string) string {
	return "Good Bye" + name
}

func main() {
	// Memanggil Function sebagai value didalam variabel
	test := getGoodBye
	test2 := getGoodBye
	fmt.Println(test(" Adrian"))
	fmt.Println(test2(" Naizar"))
}
