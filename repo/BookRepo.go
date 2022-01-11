package repo

import (
	"Library_project/other"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
)

type Book struct {
	BookId       uint16 `json:"book_id" valid:"required,numeric"`
	BookName     string `json:"book_name" valid:"required"`
	GenreId      uint16 `json:"genre_id"`
	BookGenre    string `json:"book_genre"`
	Year         uint16 `json:"year" valid:"numeric"`
	Quantity     uint16 `json:"quantity" valid:"required,numeric"`
	Available    uint16 `json:"available" valid:"required,numeric"`
	Registration string `json:"registration" valid:"required"`
	Price        uint16 `json:"book_price" valid:"required,numeric"`
	ImagePath    string `json:"image_path" valid:"required,url"`
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
	BookId     uint16 `json:"book_id"`
	BookName   string `json:"book_name"`
	AuthorId   uint16 `json:"author_id"`
	AuthorName string `json:"author_name"`
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

func GetBooksFromDBWithPages(Books *[]Book, page string) {
	db := other.ConnectDB()
	defer db.Close()
	p, _ := strconv.Atoi(page)
	get, err := db.Query(fmt.Sprintf("Select * from `books` LEFT JOIN `book_genre` ON books.book_id = book_genre.book_id LEFT JOIN `genres` ON book_genre.genre_id = genres.genre_id LIMIT 5 OFFSET %d ", (p-1)*5))

	other.CheckErr(err)

	for get.Next() {
		var genre Genre
		var book Book
		err = get.Scan(&book.BookId, &book.BookName, &book.Year, &book.Quantity, &book.Available, &book.Registration, &book.Price, &book.ImagePath, &book.BookGenre, &book.GenreId, &genre.GenreId, &genre.GenreName)
		book.GenreId = genre.GenreId
		book.BookGenre = genre.GenreName
		other.CheckErr(err)
		*Books = append(*Books, book)
	}
}

func GetBooksWithAuthorsFromDBWithPages(Books *[]BooksWithAuthors, page string) {
	db := other.ConnectDB()
	defer db.Close()
	p, _ := strconv.Atoi(page)
	get, err := db.Query(fmt.Sprintf("Select * from `books` LEFT JOIN `book_authors` ON books.book_id = book_authors.book_id LEFT JOIN `authors` ON book_authors.author_id = authors.author_id LIMIT 5 OFFSET %d ", (p-1)*5))

	other.CheckErr(err)

	for get.Next() {
		var book Book
		var author Author
		var bookAuthors BooksWithAuthors
		err = get.Scan(&book.BookId, &book.BookName, &book.Year, &book.Quantity, &book.Available, &book.Registration, &book.Price, &book.ImagePath, &book.BookGenre, &book.GenreId, &author.AuthorId, &author.AuthorName, &author.AuthorImage)
		bookAuthors.BookId = book.BookId
		bookAuthors.BookName = book.BookName
		bookAuthors.AuthorId = author.AuthorId
		bookAuthors.AuthorName = author.AuthorName
		other.CheckErr(err)
		*Books = append(*Books, bookAuthors)
	}

}
