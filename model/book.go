package model

var CurrentId int = 0

type Book struct {
	Id    int
	Title string
	Year  int
	Price float32
}

var BookList = []Book{}