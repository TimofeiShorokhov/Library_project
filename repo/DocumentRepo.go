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
