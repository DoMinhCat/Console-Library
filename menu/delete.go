package menu

import (
	"fmt"
	"library_console/model"
)

func Delete() {
	var DelId int
	var DelIndex int 
	var DelConfirm string
	var exist bool = false

	fmt.Print("ID du livre à supprimer : ")
	fmt.Scanln(&DelId)

	

	for i, book := range(model.BookList) {
		if DelId == book.Id{
			exist = true
			DelIndex = i
			break
		} else {
			fmt.Println("Livre avec ID", DelId, "introuvable !")
		}
	}

	if exist{
		fmt.Println("Etes vous sur de vouloir supprimer ce livre: ", DelId, " - ", model.BookList[DelIndex].Title, " ? (Y/N)")
		fmt.Scanln(&DelConfirm)

		if DelConfirm == "N" || DelConfirm == "n" {
			fmt.Println("Suppression du livre:", DelId, "-", model.BookList[DelIndex].Title, " a été annulée !")
		} else {
			model.BookList = append(model.BookList[0:DelIndex], model.BookList[DelIndex+1:]... )
			fmt.Println("Livre supprimé !")
		}
	}
}