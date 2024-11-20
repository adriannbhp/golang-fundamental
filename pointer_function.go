package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func ChangeCouuntyToIndonesia(address *Address) {
	address.Country = "Indonesia"
}

func main() {
	//var address *Address:= Address{}
	address := Address{}
	ChangeCouuntyToIndonesia(&address)

	fmt.Println(address)
}
