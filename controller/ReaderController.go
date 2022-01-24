package controller

import (
	"Library_project/model"
	"Library_project/other"
	"Library_project/repo"
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

var Readers []repo.Reader

func SaveReaderController(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	other.CheckErr(err)

	var reader repo.Reader

	err = json.Unmarshal(body, &reader)

	other.CheckErr(err)

	model.SaveReader(&reader)

}

func GetReadersController(w http.ResponseWriter, r *http.Request) {
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
	Readers = model.GetReadersWithPage(Readers, page, limit)
	json.NewEncoder(w).Encode(Readers)
}

func SearchReaderController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	Readers = model.GetReaders(Readers)
	for _, item := range Readers {
		if item.Name == params["name"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}
