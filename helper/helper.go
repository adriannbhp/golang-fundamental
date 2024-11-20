package helper

import "fmt"

/**
- Di bahasa pemrograman lain, biasanya ada kata knci yang bisa digunakan ntuk menentukan
  access modifier terhadap suuatu function atau variable
- Di Go-Lang, untuk menentukan access modifier, cukup dengan nama function ata variable
- Jika nama nya diawalai dengan huruf besar, maka artinya bisa diakses dari package lain,
  jika dimulai dengan huruf kecil, artinya tidak bisa diakses dari package lain
*/

var version = "1.0.0"
var Application = "golang"

func sayGoodBye(name string) string {
	return "Good Bye " + name
}

func SayHello(name string) string {
	return "Hello " + name
}

func Contoh() {
	sayGoodBye("Adrian")
	fmt.Println(version)
}
