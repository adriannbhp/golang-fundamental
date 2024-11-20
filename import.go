package main

import (
	"fmt"
	"learn-golang-basic/helper"
)

func main() {
	result := helper.SayHello("")
	fmt.Println(result)

	fmt.Println(helper.Application)
	//fmt.Println(helper.version)              // Tidak bisa diakses
	//fmt.Println(helper.sayGoodBye("Adrian")) // Tidak bisa diakses

}
