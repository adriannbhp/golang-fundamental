package main

import "fmt"

func sayHelloWithFilter(name string, filter func(string) string) {
	filteredName := filter(name) // Baris ini memanggil fungsi filter dengan name sebagai argumennya.
	fmt.Println("Hello", filteredName)
}

func spamFilter(name string) string {
	if name == "Asu" {
		return "..."
	} else {
		return name
	}
}

func goodWord(name string) string {
	if name == "Friend" {
		return "Ramah"
	} else {
		return name
	}
}

func main() {
	sayHelloWithFilter("Adrian", spamFilter)

	sayHelloWithFilter("Friend", goodWord)

	filter := spamFilter
	sayHelloWithFilter("Asu", filter)
}
