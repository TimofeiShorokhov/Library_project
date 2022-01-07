package controller

import (
	"Library_project/model"
	"Library_project/other"
	"Library_project/repo"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"time"
)

var Documents []repo.Document

func SaveDocumentController(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	other.CheckErr(err)

	var document repo.DocumentForRest
	var doc repo.Document

	err = json.Unmarshal(body, &document)

	registration := time.Now().Add(720 * time.Hour).Format("2006-01-02")
	document.Date = registration
	model.StructSwitchDoc(&doc, &document)
	if doc.BookId == "" {
		return
	} else {

		other.CheckErr(err)
		model.SaveDocument(&doc)
	}

}

func DeleteDocumentController(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	other.CheckErr(err)

	var instance repo.Instance
	Documents = model.GetDocuments(Documents)

	err = json.Unmarshal(body, &instance)
	params := mux.Vars(r)

	var penny int

	other.CheckErr(err)
	for _, i := range Documents {
		t1 := i.Date
		dt1, _ := time.Parse("2006-01-02", t1)
		t2 := time.Now()
		if i.BookId == params["instance_id"] {
			model.DeleteDocument(&i, &instance)
			if t2.After(dt1) {
				days := (t2.Sub(dt1).Hours()) / 24
				penny = int(instance.FinalPrice * (0.01) * (days))
				fmt.Fprintf(w, "'%s' вернул(а) '%s', стоимость составила: '%d'. Просрочка составила '%d' дней", i.ReaderSurname, i.BookName, penny, int(days))
			} else if t2.Before(dt1) {
				fmt.Fprintf(w, "'%s' вернул(а) '%s', стоимость составила: '%.2f'", i.ReaderSurname, i.BookName, instance.FinalPrice)
			}

		}
	}
}

func GetDocumentsController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	Documents = model.GetDocuments(Documents)
	json.NewEncoder(w).Encode(Documents)
}
