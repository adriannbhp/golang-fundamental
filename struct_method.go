package main

import "fmt"

/*
*
Struct Method
  - Struct adalah tipe data seperti tipe data lainnya, dia bisa digunakan sebagai parameter untuk
    function
  - Namun jika kita ingin menambahkan method ke dalam structs, sehingga seakan-akan sebuah struct
    memiliki function
  - Method adalah function
*/

type Username struct {
	FirstName, MiddleName, LastName string
}

func (username Username) sayHello(name string) {
	fmt.Println("Hello " + name + " My Name is " + username.FirstName)
}

func main() {

	player1 := Username{
		FirstName:  "Adrian",
		MiddleName: "Bimo",
		LastName:   "Hernawan Pratama",
	}
	fmt.Println(player1.FirstName)
	fmt.Println(player1.MiddleName)
	fmt.Println(player1.LastName)
	fmt.Println(player1)

	player1.sayHello("Naizar")
}
