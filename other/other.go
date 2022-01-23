package other

import (
	"database/sql"
	"encoding/json"
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
func RespondWithError(w http.ResponseWriter, code int, message string) {
	RespondWithJSON(w, code, map[string]string{"error": message})
}

func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func ConnectDB() *sql.DB {

	err := os.Setenv("DATABASE", "mysql:@tcp(127.0.0.1:3306)/library")
	CheckErr(err)

	database := os.Getenv("DATABASE")

	err1 := os.Setenv("DRIVER", "mysql")
	CheckErr(err1)

	driver := os.Getenv("DRIVER")

	db, err := sql.Open(driver, database)
	CheckErr(err)
	return db
}
