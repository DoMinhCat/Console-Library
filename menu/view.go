package menu

import (
	"fmt"
	"library_console/model"
)

func View() {
	if len(model.BookList) > 0 {
		for i :=0; i<len(model.BookList); i++ {
			fmt.Println(model.BookList[i].Id, "- Titre:", model.BookList[i].Title, "- Parution:", model.BookList[i].Year, "- Prix:", model.BookList[i].Price)
		}
	} else{
		fmt.Println("Aucun livre en ce moment !")
	}
}