package menu

import (
	"encoding/csv"
	"fmt"
	"library_console/model"
	"os"
	"strconv"
)

func ExportCSV() {
	// Create file with os.Create()
	file, err := os.Create("books.csv")
	if err != nil{
		fmt.Println("Error create book.csv : ", err)
		return
	}
	defer file.Close() // defer = always execute before exit, ensure file is closed

	writer := csv.NewWriter(file)
	defer writer.Flush() // ensure everything is written into file

	// Write header
	writer.Write([]string{"ID", "Titre", "Parution", "Prix"})

	// Write each book info
	for _, book := range(model.BookList){
		writer.Write([]string{strconv.Itoa(book.Id), book.Title, strconv.Itoa(book.Year), strconv.FormatFloat(float64(book.Price),'f',-1, 64)})
	}
	fmt.Println(("Tous les livres sont exportés en book.csv !"))
}