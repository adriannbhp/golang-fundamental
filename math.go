package main

import "fmt"

func main() {
	var a = 5
	var b = 5
	var c = 3
	var e = 2
	var f = a + b - c*e
	fmt.Println(f)

	/**
	Augmented Assignment
	Operasi Matematika				Augmented Assignments
	- a = a + 10					a += 10
	- a = a - 10					a -= 10
	- a = a * 10					a *= 10
	- a = a / 10					a /= 10
	- a = a % 10					a %= 10
	*/

	var i = 10
	i += 10        // i = i + 10
	fmt.Println(i) // 20

	i += 5         // i = i + 5
	fmt.Println(i) // 25

	/**
	Unary Operator			Keterangan
	- ++					a = a + 1
	- --					a = a - 1
	- -						Negative
	- +						Positive
	- !						Boolean Kebalikan
	*/

	var j = 1
	j++            // j = j + 1
	fmt.Println(j) // 2

	j++            // j = j + 1
	fmt.Println(j) // 3

	j-- // j = j - 1
	fmt.Println(j)

	j-- // j = j - 1
	fmt.Println(j)
}
