package controller

import (
	"Library_project/model"
	"Library_project/other"
	"Library_project/repo"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

var Books []repo.Book
var BooksW []repo.BooksWithAuthors

func SaveBookController(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		other.RespondWithJSON(w, 500, "Body read error")
		return
	}

	var bookForRest model.BookForRest
	var book repo.Book
	var bookGenre repo.BookGenre
	var bookAuthor repo.BookAuthor

	err = json.Unmarshal(body, &bookForRest)

	if err != nil {
		other.RespondWithJSON(w, 400, "Bad request")
		return
	}

	registration := time.Now().Format("2006-01-02")
	bookForRest.Registration = registration
	model.SaveBook(model.StructSwitchBook(&book, &bookForRest, &bookGenre, &bookAuthor))
}

func GetBooksController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	Books = model.GetBooks(Books)
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		other.RespondWithJSON(w, 500, "Body read error")
		return
	}
	page := r.URL.Query().Get("page")
	limit := r.URL.Query().Get("limit")
	err = json.Unmarshal(body, page)
	if page == "" {
		page = "1"
	}
	if limit == "" {
		limit = "1"
	}
	Books = model.GetBooksWithPage(Books, page, limit)
	json.NewEncoder(w).Encode(Books)
}

func GetBooksWithAuthorsController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		other.RespondWithJSON(w, 500, "Body read error")
		return
	}
	page := r.URL.Query().Get("page")
	limit := r.URL.Query().Get("limit")
	err = json.Unmarshal(body, page)
	if page == "" {
		page = "1"
	}
	if limit == "" {
		limit = "1"
	}
	BooksW = model.GetBooksWithAuthorsWithPage(BooksW, page, limit)
	json.NewEncoder(w).Encode(BooksW)
}

func RenderFileController(w http.ResponseWriter, r *http.Request) {
	image := r.URL.Query().Get("image")
	filename := fmt.Sprintf("D:/img/%s.jpg", image)
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		other.RespondWithJSON(w, 404, "Image not found")
	} else {
		w.Write(file)
	}
}
