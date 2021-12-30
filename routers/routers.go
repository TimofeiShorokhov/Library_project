package routers

import (
	"Library_project/controller"
	"Library_project/repo"
	"github.com/gorilla/mux"
	"net/http"
)

func Routers() {
	r := mux.NewRouter()
	repo.ConnectDB()

	http.Handle("/images/book_img", http.StripPrefix("/images/", http.FileServer(http.Dir("images"))))
	http.Handle("/images/author_img/", http.StripPrefix("/authorImage/", http.FileServer(http.Dir("authorImage"))))

	r.HandleFunc("/books/", controller.SaveBookController).Methods("POST")
	r.HandleFunc("/books", controller.GetBooksController).Methods("GET")

	http.Handle("/", r)
	http.ListenAndServe(":8080", r)
}
