package other

import (
	"database/sql"
	"io"
	"net/http"
	"os"
)

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

func ConnectDB() *sql.DB {
	db, err := sql.Open("mysql", "mysql:@tcp(127.0.0.1:3306)/library")
	CheckErr(err)
	return db
}
