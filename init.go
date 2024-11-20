package main

import (
	"fmt"
	"learn-golang-basic/database"
	_ "learn-golang-basic/internal" // Blank Identifier
)

func main() {
	fmt.Println(database.GetDatabase())
}
