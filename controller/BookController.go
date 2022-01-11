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
var BooksW []repo.BooksWithAuthors

func SaveBookController(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	other.CheckErr(err)

	var bookForRest model.BookForRest
	var book repo.Book
	var bookGenre repo.BookGenre
	var bookAuthor repo.BookAuthor

	err = json.Unmarshal(body, &bookForRest)

	other.CheckErr(err)

	registration := time.Now().Format("2006-01-02")
	bookForRest.Registration = registration
	model.SaveBook(model.StructSwitchBook(&book, &bookForRest, &bookGenre, &bookAuthor))
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

func GetBooksWithAuthorsController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	body, err := ioutil.ReadAll(r.Body)
	other.CheckErr(err)
	page := r.URL.Query().Get("page")
	err = json.Unmarshal(body, page)
	if page == "" {
		page = "1"
	}
	BooksW := model.GetBooksWithAuthorsWithPage(BooksW, page)
	json.NewEncoder(w).Encode(BooksW)
}
