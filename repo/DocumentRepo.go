package repo

import (
	"Library_project/other"
	"fmt"
)

type Document struct {
	DocId         uint16  `json:"doc_id"`
	ReaderSurname string  `json:"reader_surname"`
	BookName      string  `json:"book"`
	Date          string  `json:"date"`
	Price         float64 `json:"price"`
	QuantityBook  uint16  `json:"quant"`
}

func SaveDocumentInDB(document Document) {
	db := other.ConnectDB()
	defer db.Close()

	ins, err := db.Query(fmt.Sprintf("INSERT INTO `documents` (`reader_surname`,`book`,`date`,`price`, `quant`) VALUES ('%s','%s','%s','%f','%d')", document.ReaderSurname, document.BookName, document.Date, document.Price, document.QuantityBook))
	other.CheckErr(err)
	defer ins.Close()
}

func DeleteDocumentInDb(surname string) {
	db := other.ConnectDB()
	defer db.Close()

	updDoc := db.QueryRow("DELETE  from `documents` where reader_surname = ?", surname)
	updDoc.Err()
}

func GetDocumentsFromDB(Documents *[]Document) {
	db := other.ConnectDB()
	defer db.Close()

	get, err := db.Query("Select * from `documents`")
	other.CheckErr(err)

	for get.Next() {
		var document Document
		err = get.Scan(&document.DocId, &document.ReaderSurname, &document.BookName, &document.Date, &document.Price, &document.QuantityBook)
		other.CheckErr(err)
		*Documents = append(*Documents, document)
	}
}
