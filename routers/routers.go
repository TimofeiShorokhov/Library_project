package routers

import (
	"Library_project/controller"
	"Library_project/other"
	"github.com/gorilla/mux"
	"net/http"
)

func Routers() {
	r := mux.NewRouter()
	other.ConnectDB()

	http.Handle("/images/book_img", http.StripPrefix("/images/", http.FileServer(http.Dir("images"))))
	http.Handle("/images/author_img/", http.StripPrefix("/authorImage/", http.FileServer(http.Dir("authorImage"))))

	r.HandleFunc("/books/", controller.SaveBookController).Methods("POST")
	r.HandleFunc("/books", controller.GetBooksController).Methods("GET")

	r.HandleFunc("/readers/", controller.SaveReaderController).Methods("POST")
	r.HandleFunc("/readers", controller.GetReadersController).Methods("GET")

	http.Handle("/", r)
	http.ListenAndServe(":8080", r)
}
