package controller

import (
	"Library_project/model"
	"Library_project/other"
	"Library_project/repo"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

var Authors []repo.Author

func SaveAuthorController(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	other.CheckErr(err)

	var author repo.Author

	err = json.Unmarshal(body, &author)

	other.CheckErr(err)

	model.SaveAuthor(&author)

}

func GetAuthorsController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	Authors = model.GetAuthors(Authors)
	json.NewEncoder(w).Encode(Authors)
}
