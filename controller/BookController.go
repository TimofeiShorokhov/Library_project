package controller

import (
	"Library_project/model"
	"Library_project/other"
	"Library_project/repo"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

var Books []repo.Book

func SaveBookController(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	other.CheckErr(err)

	var bookForRest repo.BookForRest
	var book repo.Book

	err = json.Unmarshal(body, &bookForRest)

	other.CheckErr(err)

	registration := time.Now().Format("2006-01-02")
	bookForRest.Registration = registration
	model.StructSwitchBook(&book, &bookForRest)
	model.SaveBook(&book)
}

func GetBooksController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	Books = model.GetBooks(Books)
	body, err := ioutil.ReadAll(r.Body)
	other.CheckErr(err)
	page := r.URL.Query().Get("page")
	err = json.Unmarshal(body, page)
	if page == "" {
		page = "1"
	}
	Books = model.GetBooksWithPage(Books, page)
	json.NewEncoder(w).Encode(Books)
}
