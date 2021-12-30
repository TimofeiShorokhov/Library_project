package model

import (
	"Library_project/repo"
	"log"
)

type Document struct {
	DocId         uint16  `json:"doc_id"`
	ReaderSurname string  `json:"reader_surname"`
	BookName      string  `json:"book"`
	Date          string  `json:"date"`
	Price         float64 `json:"price"`
	QuantityBook  uint16  `json:"quant"`
}

func SaveDocument(document *repo.Document) {
	var Readers []repo.Reader
	Readers = GetReaders(Readers)
	if document.ReaderSurname == "" || document.BookName == "" || document.QuantityBook > 1 {
		log.Println("Пустые поля")
	} else {
		for _, item := range Readers {
			if item.Surname == document.ReaderSurname {
				if item.Debt > 0 {
					log.Println("Необходимо погасить долг прежде, чем брать новую книгу")
				} else {
					repo.SaveDocumentInDB(*document)
					repo.DecreaseBookAvailableInDB(document.BookName)
					repo.IncreaseReaderDebtInDb(document.ReaderSurname)
				}
			}
		}
	}
}

func DeleteDocument(document *repo.Document) {
	repo.DeleteDocumentInDb(document.ReaderSurname)
	repo.IncreaseBookAvailableInDB(document.BookName)
	repo.DecreaseReaderDebtInDb(document.ReaderSurname)
}

func GetDocuments(Documents []repo.Document) []repo.Document {
	Documents = []repo.Document{}
	repo.GetDocumentsFromDB(&Documents)
	return Documents
}
