package model

import (
	"Library_project/other"
	"Library_project/repo"
	"bytes"
	"fmt"
	"log"
	"strings"
)

func StructSwitchBook(book *repo.Book, bookForRest *repo.BookForRest) {
	book.BookId = bookForRest.BookId
	book.BookName = bookForRest.BookName
	book.Year = bookForRest.Year
	book.Quantity = bookForRest.Quantity
	book.Available = bookForRest.Available
	book.Registration = bookForRest.Registration
	book.Price = bookForRest.Price
	book.ImagePath = bookForRest.ImagePath
	CheckAndChangeInfoBook(book, bookForRest)
}

func CheckAndChangeInfoBook(book *repo.Book, bookForRest *repo.BookForRest) {
	var Genres []repo.Genre
	Genres = GetGenres(Genres)
	bufferGenr := bytes.Buffer{}
	genr := bookForRest.GenreId

	var Authors []repo.Author
	Authors = GetAuthors(Authors)
	bufferAuth := bytes.Buffer{}
	aut := bookForRest.AuthorId

	for _, i := range Genres {
		for _, j := range genr {
			if i.GenreName == j {
				bufferGenr.WriteString(j + ", ")
			}
		}
	}
	for _, i := range Authors {
		for _, j := range aut {
			if i.AuthorName == j {
				bufferAuth.WriteString(j + ", ")
			}
		}
	}
	a := bufferAuth.String()
	b := a[:len(a)-2]
	book.AuthorId = b

	c := bufferGenr.String()
	h := c[:len(c)-2]
	book.GenreId = h
}

func SaveBook(book *repo.Book) {

	if book.BookName == "" || book.Year <= 0 || book.Quantity <= 0 || book.Available > book.Quantity || book.Available <= 0 || book.Price <= 0 || book.ImagePath == "" {
		log.Println("Пустые поля")
		log.Println(book)
	} else {
		image := book.ImagePath
		out := strings.ReplaceAll(book.BookName, " ", "_")
		filepath := fmt.Sprintf("./images/book_img/%s.jpg", out)
		other.DownloadFile(filepath, image)
		book.ImagePath = filepath
		repo.SaveBookInDB(*book)
		repo.SaveInstanceInDB(*book)
	}
}

func GetBooks(Books []repo.Book) []repo.Book {
	Books = []repo.Book{}
	repo.GetBooksFromDB(&Books)
	return Books
}

func GetBooksWithPage(Books []repo.Book, page string) []repo.Book {
	Books = []repo.Book{}
	repo.GetBooksFromDBWithPages(&Books, page)
	return Books
}
