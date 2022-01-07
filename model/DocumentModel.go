package model

import (
	"Library_project/repo"
	"log"
	"strconv"
	"strings"
	"time"
)

func Contains(strs []string, str string) bool {
	for _, s := range strs {
		if s == str {
			return true
		}
	}
	return false
}

func Dedup(input string) bool {
	unique := []string{}

	words := strings.Split(input, " ")
	for _, word := range words {
		// If we alredy have this word, skip.
		if Contains(unique, word) {
			return false
		}

		unique = append(unique, word)
	}

	return true
}

func Discount(doc *repo.Document) {
	var Readers []repo.Reader
	Readers = GetReaders(Readers)

	for _, i := range Readers {
		if i.Surname == doc.ReaderSurname {
			if i.Debt == 2 || i.Debt == 3 {
				doc.Price -= doc.Price * 10 / 100
			} else if i.Debt >= 4 {
				doc.Price -= doc.Price * 15 / 100
			}
		}
	}
}

func StructSwitchDoc(doc *repo.Document, docs *repo.DocumentForRest) {
	doc.ReaderSurname = docs.ReaderSurname
	doc.DocId = docs.DocId
	doc.Date = docs.Date

	CheckAndChangeInfoDoc(docs, doc)
}

func CheckAndChangeInfoDoc(docs *repo.DocumentForRest, doc *repo.Document) {
	var Books []repo.Instance
	Books = GetInstances(Books)

	for _, v := range Books {
		if v.InstanceId == strconv.Itoa(int(docs.BookId)) {
			doc.BookId = v.InstanceId
			doc.BookName = v.InstanceName
			doc.Price = float64(v.InstancePrice)
			doc.QuantityBook = 1
		}
	}

}

func SaveDocument(document *repo.Document) {
	var Readers []repo.Reader
	var Books []repo.Book
	var Documents []repo.Document

	Documents = GetDocuments(Documents)
	Readers = GetReaders(Readers)
	Books = GetBooks(Books)

	for _, i := range Documents {
		if i.ReaderSurname == document.ReaderSurname {
			if i.BookName == document.BookName {
				log.Println("Только 1 экзмепляр одной книги можно взять")
				return
			}
		}
	}

	if document.ReaderSurname == "" {
		log.Println("Пустые поля")
	} else if document.QuantityBook > 5 {
		log.Println("Нельзя брать более 5 книг")
	} else {
		for _, item := range Readers {
			if item.Surname == document.ReaderSurname {
				if item.Debt > 5 {
					log.Println("Необходимо погасить долг прежде, чем брать новую книгу")
				} else {
					Discount(document)
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

func DeleteDocument(document *repo.Document, instance *repo.Instance) {
	var Books []repo.Book
	var Documents []repo.Document
	var Instances []repo.Instance

	Documents = GetDocuments(Documents)
	Books = GetBooks(Books)
	Instances = GetInstances(Instances)

	for _, i := range Documents {
		for _, v := range Instances {
			if i.BookId == v.InstanceId {
				registration := time.Now().Format("2006-01-02")
				instance.ReturnDate = registration
				repo.UpdateInstancesInDB(*instance)
			}
		}
	}

	for _, i := range Books {
		for _, v := range Instances {
			if document.BookId == v.InstanceId {
				if document.BookName == i.BookName {
					repo.IncreaseBookAvailableInDB(i.BookName)
				}
			}
		}
	}
	repo.DeleteDocumentInDb(instance.InstanceId)
	repo.DecreaseReaderDebtInDb(document.ReaderSurname)
}

func GetDocuments(Documents []repo.Document) []repo.Document {
	Documents = []repo.Document{}
	repo.GetDocumentsFromDB(&Documents)
	return Documents
}
