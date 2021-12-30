package repo

import (
	"Library_project/other"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type Book struct {
	BookId       string `json:"book_id"`
	BookName     string `json:"book_name"`
	GenreId      string `json:"book_genre_id"`
	AuthorId     string `json:"book_author_id"`
	Year         uint16 `json:"year"`
	Quantity     uint16 `json:"quantity"`
	Available    uint16 `json:"available"`
	Registration string `json:"registration"`
	Price        uint16 `json:"book_price"`
	ImagePath    string `json:"image_path"`
}

func SaveBookInDB(book Book) {
	db := other.ConnectDB()
	defer db.Close()

	ins, err := db.Query(fmt.Sprintf("INSERT INTO `books` (`book_name`,`book_genre_id`,`book_author_id`,`year`,`quantity`, `registration`, `book_price`, `available`,`image_path`) VALUES ('%s','%s','%s','%d','%d','%s','%d','%d','%s')", book.BookName, book.GenreId, book.AuthorId, book.Year, book.Quantity, book.Registration, book.Price, book.Available, book.ImagePath))
	other.CheckErr(err)
	defer ins.Close()
}

func IncreaseBookAvailableInDB(bookName string) {
	db := other.ConnectDB()
	defer db.Close()

	updBook := db.QueryRow("UPDATE `books` SET available=available-1 where book_name = ?", bookName)
	updBook.Err()
}

func GetBooksFromDB(Books *[]Book) {
	db := other.ConnectDB()
	defer db.Close()

	get, err := db.Query("Select * from `books` order by book_name")
	other.CheckErr(err)

	for get.Next() {
		var book Book
		err = get.Scan(&book.BookId, &book.BookName, &book.GenreId, &book.AuthorId, &book.Year, &book.Quantity, &book.Available, &book.Registration, &book.Price, &book.ImagePath)
		other.CheckErr(err)
		*Books = append(*Books, book)
	}
}
