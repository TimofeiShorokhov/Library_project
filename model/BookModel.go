package model

import (
	"Library_project/other"
	"Library_project/repo"
	"fmt"
	"github.com/asaskevich/govalidator"
	"log"
	"strings"
)

type BookForRest struct {
	BookId       uint16   `json:"book_id" `
	BookName     string   `json:"book_name" `
	GenreId      []uint16 `json:"genre_id"`
	AuthorId     []uint16 `json:"author_id"`
	Year         uint16   `json:"year" `
	Quantity     uint16   `json:"quantity"`
	Available    uint16   `json:"available"`
	Registration string   `json:"registration"`
	Price        uint16   `json:"book_price"`
	ImagePath    string   `json:"image_path"`
}

func StructSwitchBook(book *repo.Book, bookForRest *BookForRest, bookGenre *repo.BookGenre, bookAuthor *repo.BookAuthor) (*[]repo.BookAuthor, *[]repo.BookGenre, *repo.Book) {
	var bokGenre []repo.BookGenre
	var bokAuthor []repo.BookAuthor

	book.BookId = bookForRest.BookId
	book.BookName = bookForRest.BookName
	book.Year = bookForRest.Year
	book.Quantity = bookForRest.Quantity
	book.Available = bookForRest.Available
	book.Registration = bookForRest.Registration
	book.Price = bookForRest.Price
	book.ImagePath = bookForRest.ImagePath

	for _, h := range bookForRest.GenreId {
		bookGenre.BookGenre = h
		bookGenre.BookId = bookForRest.BookId
		bokGenre = append(bokGenre, *bookGenre)
	}

	for _, c := range bookForRest.AuthorId {
		bookAuthor.BookId = bookForRest.BookId
		bookAuthor.BookAuthor = c
		bokAuthor = append(bokAuthor, *bookAuthor)
	}
	res, err := govalidator.ValidateStruct(bookGenre)
	other.CheckErr(err)

	res1, err1 := govalidator.ValidateStruct(bookAuthor)
	other.CheckErr(err1)

	if res == true || res1 == true {
		return &bokAuthor, &bokGenre, book
	}
	return nil, nil, nil
}

func SaveBook(bookAuthor *[]repo.BookAuthor, bookGenre *[]repo.BookGenre, book *repo.Book) {
	var oneBookGenre repo.BookGenre
	var oneBookAuthor repo.BookAuthor

	result, err := govalidator.ValidateStruct(book)
	other.CheckErr(err)

	if result != true {
		log.Println(err)
	} else {
		image := book.ImagePath
		out := strings.ReplaceAll(book.BookName, " ", "_")
		filepath := fmt.Sprintf("./images/book_img/%s.jpg", out)
		other.DownloadFile(filepath, image)
		book.ImagePath = filepath
		repo.SaveBookInDB(*book)

		for _, h := range *bookGenre {
			oneBookGenre = h
			repo.SaveBookGenreInDB(oneBookGenre)
		}

		for _, c := range *bookAuthor {
			oneBookAuthor = c
			repo.SaveBookAuthorInDB(oneBookAuthor)
		}

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

func GetBooksWithAuthorsWithPage(Books []repo.BooksWithAuthors, page string) []repo.BooksWithAuthors {
	Books = []repo.BooksWithAuthors{}

	repo.GetBooksWithAuthorsFromDBWithPages(&Books, page)
	return Books
}
