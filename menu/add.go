package menu

import (
	"fmt"
	"library_console/model"
)

func Add() {
	var newTitle string
	var newYear int
	var newPrice float32

	fmt.Print("Nom du livre : ")
	fmt.Scanln(&newTitle)
	fmt.Print("Année de sortie du livre : ")
	fmt.Scanln(&newYear)
	fmt.Print("Prix du livre : ")
	fmt.Scanln(&newPrice)

	var newBook model.Book = model.Book{Id: model.CurrentId+1,Title: newTitle, Year: newYear, Price: newPrice}
	model.BookList = append(model.BookList, newBook)
	model.CurrentId++

	fmt.Printf("Livre \"%v\" a été ajouté !", newTitle)
	fmt.Println("")
}