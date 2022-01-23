package controller

import (
	"Library_project/model"
	"Library_project/other"
	"Library_project/repo"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

var Instances []repo.Instance

func GetInstancesController(w http.ResponseWriter, r *http.Request) {
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

	Instances = model.GetInstancesWithPage(Instances, page, limit)
	json.NewEncoder(w).Encode(Instances)
}
