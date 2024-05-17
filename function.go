package main

import "fmt"

func sayHello() {
	fmt.Println("Hello")
}

// Bisa memanggil function berkali kali
func main() {
	sayHello()
	sayHello()
}
