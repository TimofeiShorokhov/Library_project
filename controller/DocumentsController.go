package controller

import (
	"Library_project/model"
	"Library_project/other"
	"Library_project/repo"
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var Documents []repo.Document

func SaveDocumentController(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	other.CheckErr(err)

	var document repo.Document

	err = json.Unmarshal(body, &document)

	registration := time.Now().Add(720 * time.Hour).Format("2006-01-02")
	document.Date = registration

	other.CheckErr(err)
	model.SaveDocument(&document)
}

func DeleteDocumentController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	Documents = model.GetDocuments(Documents)
	if repo.SearchReaderInDb(params["reader_surname"]) == true {
		for _, item := range Documents {
			if item.ReaderSurname == params["reader_surname"] {
				model.DeleteDocument(&item)
			}
		}
	} else {
		log.Println("Error")
		return
	}
}

func GetDocumentsController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	Documents = model.GetDocuments(Documents)
	json.NewEncoder(w).Encode(Documents)
}
