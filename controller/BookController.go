package controller

import (
	"Library_project/model"
	"Library_project/repo"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

var Books []repo.Book

func SaveBookController(w http.ResponseWriter, r *http.Request){
	var book repo.Book

	body, err := ioutil.ReadAll(r.Body)
	repo.CheckErr(err)

	keyVal := make(map[string]string)

	json.Unmarshal(body, &keyVal)
	registration := time.Now().Format("2006-01-02")
	book.BookName = keyVal["book_name"]
	book.GenreId = keyVal["book_genre_id"]
	book.AuthorId = keyVal["book_author_id"]
	year :=  keyVal["year"]
	quantity := keyVal["quantity"]
	book.Registration = registration
	price := keyVal["book_price"]
	available := keyVal["available"]
	book.ImagePath = keyVal["image_path"]

	yearInDb, _ := strconv.Atoi(year)
	book.Year = uint16(yearInDb)

	quantityInDb, _ := strconv.Atoi(quantity)
	book.Quantity = uint16(quantityInDb)

	priceInDb, _ := strconv.Atoi(price)
	book.Price = uint16(priceInDb)

	availableInDb, _ := strconv.Atoi(available)
	book.Available = uint16(availableInDb)

	model.SaveBook(&book)
}

func GetBooksController(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	Books = model.GetBooks(Books)
	json.NewEncoder(w).Encode(Books)
}
