package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue // Melanjutkan ke iterasi selanjutnya
		}
		fmt.Println("Perualangan ke", i)
	}
}
