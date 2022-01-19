package controller

import (
	"Library_project/model"
	"Library_project/other"
	"Library_project/repo"
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"strconv"
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
	other.CheckErr(err)
	model.SaveDocument(&doc)

}

func DeleteDocumentController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
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
		str, _ := strconv.Atoi(params["instance_id"])
		if i.BookId == uint16(str) {
			instance.InstanceName = i.BookName
			model.DeleteDocument(&i, &instance)
			if t2.After(dt1) {
				days := (t2.Sub(dt1).Hours()) / 24
				penny = int(instance.FinalPrice * (0.01) * (days))
				instance.FinalPrice = float64(penny) + instance.FinalPrice
				json.NewEncoder(w).Encode(instance)
			} else if t2.Before(dt1) {
				json.NewEncoder(w).Encode(instance)
			}

		}
	}
}

func GetDocumentsController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	Documents = model.GetDocuments(Documents)
	json.NewEncoder(w).Encode(Documents)
}

func SearchDocumentController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	Documents = model.GetDocuments(Documents)
	for _, item := range Documents {
		if item.ReaderSurname == params["surname"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}
