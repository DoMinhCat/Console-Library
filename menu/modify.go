package menu

import (
	"fmt"
	"library_console/model"
)

func Modify() {
	var ModId int
	var ModIndex int 
	var ModContinue bool = true
	var choiceMod int
	var exist bool = false
	
	var newTitle string
	var newYear int
	var newPrice float32

	fmt.Print("ID du livre à modifié : ")
	fmt.Scanln(&ModId)

	for i, book := range(model.BookList) {
		if ModId == book.Id{
			exist = true
			ModIndex = i
			break
		} else {
			fmt.Println("Livre avec ID", ModId, "introuvable !")
		}
	}

	if exist {
		for ModContinue {
			fmt.Println(`Que voulez vous modifier ?
				1 - Titre
				2 - Parution
				3 - Prix
				4 - Terminer la modification`)
			fmt.Scanln(&choiceMod)
			
			switch choiceMod{
			case 1:
				fmt.Print("Nouveau titre: ")
				fmt.Scanln(&newTitle)

				model.BookList[ModIndex].Title = newTitle
				fmt.Println("Titre modifié !")
			case 2:
				fmt.Print("Nouvelle parution: ")
				fmt.Scanln(&newYear)

				model.BookList[ModIndex].Year = newYear
				fmt.Println("Parution modifiée !")
			case 3:
				fmt.Print("Nouveau prix: ")
				fmt.Scanln(&newPrice)

				model.BookList[ModIndex].Price = newPrice
				fmt.Println("Prix modifié !")
			case 4:
				fmt.Println("Modification terminé !")
				ModContinue = false
			default:
				fmt.Println("Invalide. Veuillez choisir une de ces actions ci-dessus !")
			}
		}
	} 
}