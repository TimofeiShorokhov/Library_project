package model

import (
	"Library_project/other"
	"Library_project/repo"
	"fmt"
	"log"
)

func SaveBook(book *repo.Book) {

	if book.BookName == "" || book.GenreId == "" || book.AuthorId == "" || book.Year <= 0 || book.Quantity <= 0 || book.Available > book.Quantity || book.Available <= 0 || book.Price <= 0 || book.ImagePath == "" {
		log.Println("Пустые поля")
	} else {
		image := book.ImagePath
		filepath := fmt.Sprintf("./images/book_img/%s.jpg", book.BookName)
		other.DownloadFile(filepath, image)
		book.ImagePath = filepath
		repo.SaveBookInDB(*book)
	}
}

func GetBooks(Books []repo.Book) []repo.Book {
	Books = []repo.Book{}
	repo.GetBooksFromDB(&Books)
	return Books
}
