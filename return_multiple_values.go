package main

import "fmt"

// Returning Multiple Values
func getFullName() (string, string) {
	return "Adrian", "Bimo"
}

func main() {
	//firstName, lastName := getFullName()
	//fmt.Println(firstName, lastName)

	// Menghiraukan Return Value (menggunakan _)
	firstName, _ := getFullName()
	fmt.Println(firstName)
}
