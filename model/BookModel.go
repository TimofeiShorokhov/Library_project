package model

import (
	"Library_project/repo"
	"fmt"
	"log"
)

type Book struct{
	BookId   string `json:"book_id"`
	BookName string `json:"book_name"`
	GenreId    string `json:"book_genre_id"`
	AuthorId string `json:"book_author_id"`
	Year uint16 `json:"year"`
	Quantity uint16 `json:"quantity"`
	Available uint16 `json:"available"`
	Registration string `json:"registration"`
	Price uint16 `json:"book_price"`
	ImagePath string `json:"image_path"`
}

func SaveBook(book *repo.Book) {


	if book.BookName == "" || book.GenreId == "" || book.AuthorId == "" || book.Year <= 0 || book.Quantity <= 0 || book.Available > book.Quantity || book.Available <= 0 || book.Price <= 0 || book.ImagePath == ""{
		log.Println("Пустые поля")
	} else {
		image := book.ImagePath
		filepath := fmt.Sprintf("./img/%s.jpg",book.BookName)
		repo.DownloadFile(filepath,image)
		book.ImagePath = filepath
		repo.SaveBookInDB(*book)
	}
}

func GetBooks(Books []repo.Book) []repo.Book{
	Books = []repo.Book{}
	repo.GetBooksFromDB(&Books)
	return Books
}
