package main

import (
	"book-store/utils"
	"fmt"
)

func main() {
	fmt.Println("welcome to the bookstore")
	utils.CreateLettersList()
	fmt.Println(utils.GenerateID())
}
