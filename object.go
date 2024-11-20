package main

import "fmt"

// Mendefinisikan struct Ebook
type Ebook struct {
	ID            uint
	Title         string
	Author        string
	Publisher     string
	PublishedYear uint
	ISBN          string
	URLFile       string
}

func (e Ebook) Summary() string {
	return fmt.Sprintf("'%s' by %s (%d)", e.Title, e.Author, e.PublishedYear)
}

func main() {

	// Membuat objek Ebook baru
	myEbook := Ebook{
		ID:            1,
		Title:         "Learning Go",
		Author:        "John Doe",
		Publisher:     "Tech Books Publishing",
		PublishedYear: 2024,
		ISBN:          "231231123",
		URLFile:       "http://",
	}

	// Menampilkan detail Ebook
	fmt.Printf("Ebook Details:\n")
	fmt.Printf("ID: %d\n", myEbook.ID)
	fmt.Printf("Title: %s\n", myEbook.Title)
	fmt.Printf("Author: %s\n", myEbook.Author)
	fmt.Printf("Publisher: %s\n", myEbook.Publisher)
	fmt.Printf("Published Year: %d\n", myEbook.PublishedYear)
	fmt.Printf("ISBN: %s\n", myEbook.ISBN)
	fmt.Printf("URL: %s\n", myEbook.URLFile)
	fmt.Println(myEbook.Summary())
}
