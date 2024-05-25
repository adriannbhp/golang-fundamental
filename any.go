package main

import "fmt"

/*
*
Interface Kosong
  - Go-Lang bukanlah bahasa pemrograman yang berorientasi object
  - Biasanya dalam pemrograman berorientasi object, ada satu data parent di puncak yang bisa dianggap
    sebagai semua implementasi data yang ada di bahasa pemrograman tersebut
  - Contoh di Java ada java.lang.Object
  - Untuk menangani kasus seperti ini, di Go-Lang kita bisa menggunakan interface kosong
  - Interface kosong adalah interface yang tidak memiliki deklarasi method satupun, hal ini membuat
    secara otomatis semua tipe data akan menjadi implementasi nya
  - Interface kosong, juga memiliki alias bernama any
*/

func Ups() any {
	//return 1
	//return true
	return "Ups"
}
func main() {
	var kosong any = Ups()
	fmt.Println(kosong)
}
