package repo

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io"
	"net/http"
	"os"
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

func DownloadFile(filepath string, url string) error {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

func ConnectDB() *sql.DB{
	db, err := sql.Open("mysql", "mysql:@tcp(127.0.0.1:3306)/library")
	CheckErr(err)
	return db
}


func SaveBookInDB(book Book){
	db:= ConnectDB()
	defer db.Close()

	ins, err := db.Query(fmt.Sprintf("INSERT INTO `books` (`book_name`,`book_genre_id`,`book_author_id`,`year`,`quantity`, `registration`, `book_price`, `available`,`image_path`) VALUES ('%s','%s','%s','%d','%d','%s','%d','%d','%s')",book.BookName, book.GenreId, book.AuthorId,book.Year,book.Quantity,book.Registration,book.Price,book.Available,book.ImagePath))
	CheckErr(err)
	defer ins.Close()
}

func GetBooksFromDB(Books *[]Book){
	db:= ConnectDB()
	defer db.Close()

	get, err := db.Query("Select * from `books` order by book_name")
	CheckErr(err)

		for get.Next() {
			var book Book
			err = get.Scan(&book.BookId, &book.BookName, &book.GenreId, &book.AuthorId, &book.Year, &book.Quantity, &book.Available, &book.Registration, &book.Price, &book.ImagePath)
			CheckErr(err)
			*Books = append(*Books, book)
		}
	}


