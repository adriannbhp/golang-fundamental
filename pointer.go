package main

/*
*
Pass by Value
  - Secara default di Go-Lang semua variable itu di passing, bukan by reference
  - Artinya, jika kita mengirim sebuah variable ke dalam function, method atau variable lain,
    sebenarnya yang dikirim adalah duplikasi value nya
*/
type Address struct {
	City, Province, Country string
}

func main() {
	// Program Pass by Value

}
