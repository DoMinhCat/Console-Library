package menu

import (
	"fmt"
	"library_console/model"
)

// Bubble sort
func SortByYear() {
	var nSwap int = 1

	for nSwap != 0 {
		nSwap = 0
		for i :=0; i<len(model.BookList)-1; i++{
			if model.BookList[i].Year < model.BookList[i+1].Year{
				nSwap++
				model.BookList[i], model.BookList[i+1] = model.BookList[i+1], model.BookList[i]
			}
		}
	}
	fmt.Println("Liste de livres a été triée par année de sortie !")
}

func SortByTitle() {
	var nSwap int = 1

	for nSwap != 0 {
		nSwap = 0
		for i :=0; i<len(model.BookList)-1; i++{
			if model.BookList[i].Title > model.BookList[i+1].Title{
				nSwap++
				model.BookList[i], model.BookList[i+1] = model.BookList[i+1], model.BookList[i]
			}
		}
	}
	fmt.Println("Liste de livres a été triée par titre !")
}

