package controller

import (
	"Library_project/model"
	"Library_project/other"
	"Library_project/repo"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
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
		log.Printf("Body read error, %v", err)
		w.WriteHeader(500)
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
		fmt.Println("Can't open file: " + filename)
	} else {
		w.Write(file)
	}
}
