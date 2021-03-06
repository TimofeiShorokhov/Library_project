package repo

import (
	"Library_project/other"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
)

type Book struct {
	BookId       uint16  `json:"book_id" valid:"required,numeric"`
	BookName     string  `json:"book_name" valid:"required"`
	Genre        []Genre `json:"genre"`
	Year         uint16  `json:"year" valid:"numeric"`
	Quantity     uint16  `json:"quantity" valid:"required,numeric"`
	Available    uint16  `json:"available" valid:"required,numeric"`
	Registration string  `json:"registration" valid:"required"`
	Price        uint16  `json:"book_price" valid:"required,numeric"`
	ImagePath    string  `json:"image_path" valid:"required,url"`
}

type BookGenre struct {
	BookId    uint16 `json:"book_id" valid:"required"`
	BookGenre uint16 `json:"book_genre" valid:"required,numeric"`
}

type BookAuthor struct {
	BookId     uint16 `json:"book_id" valid:"required"`
	BookAuthor uint16 `json:"book_author" valid:"required,numeric"`
}

type BooksWithAuthors struct {
	BookId   uint16   `json:"book_id"`
	BookName string   `json:"book_name"`
	Authors  []Author `json:"authors"`
}

func SaveBookInDB(book Book) {
	db := other.ConnectDB()
	defer db.Close()

	ins, err := db.Query(fmt.Sprintf("INSERT INTO `books` (`book_id`, `book_name`, `year`, `quantity`, `available`, `registration`, `book_price`, `Image_path`) VALUES ( '%d','%s', '%d', '%d', '%d', '%s', '%d', '%s');", book.BookId, book.BookName, book.Year, book.Quantity, book.Available, book.Registration, book.Price, book.ImagePath))
	other.CheckErr(err)
	defer ins.Close()

}

func SaveBookGenreInDB(genre BookGenre) {
	db := other.ConnectDB()
	defer db.Close()

	ins, err := db.Query(fmt.Sprintf("INSERT INTO `book_genre` (`book_id`, `genre_id`) VALUES ('%d','%d');", genre.BookId, genre.BookGenre))
	other.CheckErr(err)
	defer ins.Close()

}

func SaveBookAuthorInDB(author BookAuthor) {
	db := other.ConnectDB()
	defer db.Close()

	ins, err := db.Query(fmt.Sprintf("INSERT INTO `book_authors` (`book_id`, `author_id`) VALUES ('%d','%d');", author.BookId, author.BookAuthor))
	other.CheckErr(err)
	defer ins.Close()

}

func DecreaseBookAvailableInDB(bookName string) {
	db := other.ConnectDB()
	defer db.Close()

	updBook := db.QueryRow("UPDATE `books` SET available=available-1 where book_name = ?", bookName)
	updBook.Err()
}
func IncreaseBookAvailableInDB(bookName string) {
	db := other.ConnectDB()
	defer db.Close()

	updBook := db.QueryRow("UPDATE `books` SET available=available+1 where book_name = ?", bookName)
	updBook.Err()
}

func GetBooksFromDB(Books *[]Book) {
	db := other.ConnectDB()
	defer db.Close()

	get, err := db.Query("Select * from `books` order by book_name")
	other.CheckErr(err)

	for get.Next() {
		var book Book

		err = get.Scan(&book.BookId, &book.BookName, &book.Year, &book.Quantity, &book.Available, &book.Registration, &book.Price, &book.ImagePath)
		other.CheckErr(err)

		*Books = append(*Books, book)
	}
}

func GetBooksFromDBWithPages(Books *[]Book, page string, limit string) {

	db := other.ConnectDB()
	defer db.Close()
	p, _ := strconv.Atoi(page)
	l, _ := strconv.Atoi(limit)
	pageForSql := (p - 1) * 5

	get, err := db.Query(fmt.Sprintf("Select * from `books` LIMIT %d OFFSET %d ", l, pageForSql))

	other.CheckErr(err)

	for get.Next() {
		var book Book
		err = get.Scan(&book.BookId, &book.BookName, &book.Year, &book.Quantity, &book.Available, &book.Registration, &book.Price, &book.ImagePath)
		book.Genre = SelectGenres(book.BookId)

		other.CheckErr(err)
		*Books = append(*Books, book)
	}
}

func GetBooksWithAuthorsFromDBWithPages(Books *[]BooksWithAuthors, page string, limit string) {

	db := other.ConnectDB()
	defer db.Close()
	p, _ := strconv.Atoi(page)
	l, _ := strconv.Atoi(limit)
	pageForSql := (p - 1) * l

	get, err := db.Query(fmt.Sprintf("Select books.book_id, books.book_name from `books` LIMIT %d OFFSET %d ", l, pageForSql))
	other.CheckErr(err)
	for get.Next() {
		var book Book
		var bookAuthors BooksWithAuthors
		err = get.Scan(&book.BookId, &book.BookName)
		bookAuthors.BookId = book.BookId
		bookAuthors.BookName = book.BookName
		other.CheckErr(err)
		bookAuthors.Authors = SelectAuthors(book.BookId)
		*Books = append(*Books, bookAuthors)
	}

}

func SelectAuthors(id uint16) []Author {
	var authors []Author
	db := other.ConnectDB()
	defer db.Close()

	get, err := db.Query(fmt.Sprintf("SELECT authors.author_id, author_name, author_image FROM authors JOIN book_authors ON authors.author_id = book_authors.author_id AND book_authors.book_id = %d", id))
	other.CheckErr(err)
	for get.Next() {
		var author Author
		err = get.Scan(&author.AuthorId, &author.AuthorName, &author.AuthorImage)
		authors = append(authors, author)
	}
	return authors
}

func SelectGenres(id uint16) []Genre {
	var genres []Genre
	db := other.ConnectDB()
	defer db.Close()

	get, err := db.Query(fmt.Sprintf("SELECT genres.genre_id, book_genre FROM genres JOIN book_genre ON genres.genre_id = book_genre.genre_id AND book_genre.book_id = %d", id))

	other.CheckErr(err)
	for get.Next() {
		var genre Genre
		err = get.Scan(&genre.GenreId, &genre.GenreName)
		genres = append(genres, genre)
	}
	return genres
}
