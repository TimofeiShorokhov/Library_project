package routers

import (
	"Library_project/controller"
	"Library_project/other"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

func Routers() {
	r := mux.NewRouter()
	other.ConnectDB()

	err := os.Setenv("PORT", ":8080")
	other.CheckErr(err)

	port := os.Getenv("PORT")

	http.Handle("/book_img/", http.StripPrefix("/book_img/", http.FileServer(http.Dir("./book_img"))))
	http.Handle("/images/author_img/", http.StripPrefix("/authorImage/", http.FileServer(http.Dir("authorImage"))))

	r.HandleFunc("/books/", controller.SaveBookController).Methods("POST")
	r.HandleFunc("/books", controller.GetBooksController).Methods("GET")
	r.HandleFunc("/booksW", controller.GetBooksWithAuthorsController).Methods("GET")

	r.HandleFunc("/readers/", controller.SaveReaderController).Methods("POST")
	r.HandleFunc("/readers", controller.GetReadersController).Methods("GET")
	r.HandleFunc("/search_reader/{name}", controller.SearchReaderController).Methods("GET")

	r.HandleFunc("/documents", controller.GetDocumentsController).Methods("GET")

	r.HandleFunc("/take/", controller.SaveDocumentController).Methods("POST")

	r.HandleFunc("/refund_book/{instance_id}", controller.DeleteDocumentController).Methods("POST")

	r.HandleFunc("/authors/", controller.SaveAuthorController).Methods("POST")
	r.HandleFunc("/authors", controller.GetAuthorsController).Methods("GET")

	r.HandleFunc("/instances", controller.GetInstancesController).Methods("GET")

	http.Handle("/", r)
	http.ListenAndServe(port, r)
}
