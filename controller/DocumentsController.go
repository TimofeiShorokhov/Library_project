package controller

import (
	"Library_project/model"
	"Library_project/other"
	"Library_project/repo"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

var Documents []repo.Document

func SaveDocumentController(w http.ResponseWriter, r *http.Request) {
	var document repo.Document

	body, err := ioutil.ReadAll(r.Body)
	other.CheckErr(err)

	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)
	registration := time.Now().Add(720 * time.Hour).Format("2006-01-02")
	document.ReaderSurname = keyVal["reader_surname"]
	document.BookName = keyVal["book"]
	price := keyVal["price"]
	quantity := keyVal["quant"]
	document.Date = registration

	priceInDb, _ := strconv.Atoi(price)
	document.Price = float64(priceInDb)

	quantityInDb, _ := strconv.Atoi(quantity)
	document.QuantityBook = uint16(quantityInDb)

	model.SaveDocument(&document)
}
