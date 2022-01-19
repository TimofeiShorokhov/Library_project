package controller

import (
	"Library_project/model"
	"Library_project/other"
	"Library_project/repo"
	"encoding/json"
	"fmt"
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
	page := r.URL.Query().Get("page")
	limit := r.URL.Query().Get("limit")
	if page == "" {
		page = "1"
	}
	if limit == "" {
		limit = "1"
	}
	Authors = model.GetAuthors(Authors, page, limit)
	json.NewEncoder(w).Encode(Authors)
}

func RenderAuthorFileController(w http.ResponseWriter, r *http.Request) {

	image := r.URL.Query().Get("image")
	filename := fmt.Sprintf("D:/authors/%s.jpg", image)
	fmt.Println("Read request: " + filename)
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Cann't open file: " + filename)
	} else {
		w.Write(file)
	}
}
