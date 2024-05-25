package main

import "fmt"

/*
*
Recursive Function
  - Recursive function adalah function yang memanggil function dirinya sendiri
  - Kadang dalam pekerjaan, kita sering menemui kasus dimana menggunakan recursive function
    lebih mudah dibandingkan tidak menggunakan recursive function
  - Contoh kasus yang lebih mudah diselesaikan menggunakan recursive adalah Factorial
*/

// Membuat Function Factorial For Loop
func factorialLoop(value int) int {
	result := 1
	for i := value; i > 0; i-- {
		result *= i
	}
	return result
}

// Membuata Factorial Recursive (Memanggil function dirinya sendiri)
func FactorialRecursive(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * FactorialRecursive(value-1)
	}
}

func main() {
	fmt.Println(factorialLoop(10))
	fmt.Println(FactorialRecursive(10))
}
