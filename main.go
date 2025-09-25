package main

import (
	"fmt"
	"library_console/menu"
)

func main(){
	var isEnded bool = false

	for !isEnded {
		menu.PrintMenu()
		var choice int = menu.GetChoice()
		
		switch choice{
		case 1:
			menu.View()
		case 2:
			menu.Add()
		case 3:
			menu.Modify()
		case 4:
			menu.Delete()
		case 5:
			menu.SortByYear()
		case 6:
			menu.SortByTitle()
		case 7:
			menu.ExportCSV()
		case 8:
			isEnded = true
		default:
			fmt.Println("Invalide. Veuillez choisir une de ces actions ci-dessus !")
		}
	}
	fmt.Print("Au revoir !")
}
