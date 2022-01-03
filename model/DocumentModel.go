package model

import (
	"Library_project/repo"
	"bytes"
	"log"
	"strings"
)

func Discount(doc *repo.Document) {
	if doc.QuantityBook == 2 || doc.QuantityBook == 3 {
		doc.Price -= doc.Price * 10 / 100
	} else if doc.QuantityBook >= 4 {
		doc.Price -= doc.Price * 15 / 100
	}
}

func StructSwitch(doc *repo.Document, docs *repo.DocumentForRest) {
	doc.ReaderSurname = docs.ReaderSurname
	doc.DocId = docs.DocId
	doc.Date = docs.Date
	doc.QuantityBook = docs.QuantityBook
	doc.Price = docs.Price
	CheckAndChangeInfo(docs, doc)
}

func CheckAndChangeInfo(docs *repo.DocumentForRest, doc *repo.Document) {
	var Books []repo.Book
	Books = GetBooks(Books)
	books := docs.BookName
	buffer := bytes.Buffer{}
	var price uint16
	var quant uint16
	for _, i := range Books {
		for _, j := range books {
			if i.BookName == j {
				buffer.WriteString(j + ", ")
				price += i.Price
				quant = quant + 1
			} else {
				log.Println("Something wrong")
			}
		}
	}
	doc.BookName = buffer.String()
	doc.Price = float64(price)
	doc.QuantityBook = quant
	Discount(doc)
}

func SaveDocument(document *repo.Document) {
	var Readers []repo.Reader
	var Books []repo.Book
	Readers = GetReaders(Readers)
	Books = GetBooks(Books)
	if document.ReaderSurname == "" || document.BookName == "" {
		log.Println("Пустые поля")
	} else if document.QuantityBook > 5 {
		log.Println("Нельзя брать более 5 книг")
	} else {
		for _, item := range Readers {
			if item.Surname == document.ReaderSurname {
				if item.Debt > 0 {
					log.Println("Необходимо погасить долг прежде, чем брать новую книгу")
				} else {
					repo.SaveDocumentInDB(*document)
					for _, i := range Books {
						if strings.Contains(document.BookName, i.BookName) {
							repo.DecreaseBookAvailableInDB(i.BookName)
						}
					}
					repo.IncreaseReaderDebtInDb(document.ReaderSurname, document.QuantityBook)
				}
			}
		}
	}
}

func DeleteDocument(document *repo.Document) {
	var Books []repo.Book
	Books = GetBooks(Books)
	for _, i := range Books {
		if strings.Contains(document.BookName, i.BookName) {
			repo.IncreaseBookAvailableInDB(i.BookName)
		}
	}
	repo.DeleteDocumentInDb(document.ReaderSurname)
	repo.DecreaseReaderDebtInDb(document.ReaderSurname, document.QuantityBook)
}

func GetDocuments(Documents []repo.Document) []repo.Document {
	Documents = []repo.Document{}
	repo.GetDocumentsFromDB(&Documents)
	return Documents
}
