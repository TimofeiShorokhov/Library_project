package repo

import (
	"Library_project/other"
	"fmt"
)

type Document struct {
	DocId         uint16  `json:"doc_id"`
	ReaderSurname string  `json:"reader_surname" valid:"required"`
	BookId        uint16  `json:"book_id" valid:"required,numeric"`
	BookName      string  `json:"book_name"`
	Date          string  `json:"date"`
	Price         float64 `json:"price"`
	QuantityBook  uint16  `json:"quant"`
}
type DocumentForRest struct {
	DocId         uint16  `json:"doc_id"`
	ReaderSurname string  `json:"reader_surname"`
	BookId        uint16  `json:"book_id"`
	BookName      string  `json:"book_name"`
	Date          string  `json:"date"`
	Price         float64 `json:"price"`
	QuantityBook  uint16  `json:"quant"`
}

func SaveDocumentInDB(document Document) {
	db := other.ConnectDB()
	defer db.Close()

	ins, err := db.Query(fmt.Sprintf("INSERT INTO `documents` (`reader_surname`,`book_id`,`book_name`,`date`,`price`, `quant`) VALUES ('%s','%d','%s','%s','%f','%d')", document.ReaderSurname, document.BookId, document.BookName, document.Date, document.Price, document.QuantityBook))
	other.CheckErr(err)
	defer ins.Close()
}

func DeleteDocumentInDb(id uint16) {
	db := other.ConnectDB()
	defer db.Close()

	updDoc := db.QueryRow("DELETE  from `documents` where book_id = ?", id)
	updDoc.Err()
}

func GetDocumentsFromDB(Documents *[]Document) {
	db := other.ConnectDB()
	defer db.Close()

	get, err := db.Query("Select * from `documents`")
	other.CheckErr(err)

	for get.Next() {
		var document Document
		err = get.Scan(&document.DocId, &document.ReaderSurname, &document.BookId, &document.BookName, &document.Date, &document.Price, &document.QuantityBook)
		other.CheckErr(err)
		*Documents = append(*Documents, document)
	}
}
