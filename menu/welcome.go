package menu

import "fmt"
func PrintMenu(){
	var msgMenu string = `Bienvenue à la bibliothèque, que voulez vous faire ?
			1 - Voir liste des livres
			2 - Ajouter un livre
			3 - Modifier un livre
			4 - Supprimer un livre
			5 - Trier par année de sortie (les plus récents)
			6 - Trier par titre (A à Z)
			7 - Exporter un fichier CSV des livres
			8 - Sortir de la biliothèque`	
			
	fmt.Println(msgMenu)
}

func GetChoice() int {
	var choiceMenu int

	fmt.Print("Votre choix: ")
	fmt.Scanln(&choiceMenu)
	
	return choiceMenu
}
