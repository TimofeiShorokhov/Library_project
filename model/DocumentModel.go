package model

import (
	"Library_project/other"
	"Library_project/repo"
	"github.com/asaskevich/govalidator"
	"log"
	"strings"
	"time"
)

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
		if v.InstanceId == docs.BookId {
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
	var name string

	Documents = GetDocuments(Documents)
	Readers = GetReaders(Readers)
	Books = GetBooks(Books)

	res, err := govalidator.ValidateStruct(document)
	other.CheckErr(err)

	for _, i := range Documents {
		if i.ReaderSurname == document.ReaderSurname {
			if i.BookName == document.BookName {
				log.Println("Только 1 экзмепляр одной книги можно взять")
				return
			}
		}
	}

	if res != true {
		log.Println(err)
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
							name = i.BookName

						}
					}
					repo.DecreaseBookAvailableInDB(name)
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
	var name string

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
					name = i.BookName
				}
			}
		}
	}
	repo.IncreaseBookAvailableInDB(name)
	repo.DeleteDocumentInDb(instance.InstanceId)
	repo.DecreaseReaderDebtInDb(document.ReaderSurname)
}

func GetDocuments(Documents []repo.Document) []repo.Document {
	Documents = []repo.Document{}
	repo.GetDocumentsFromDB(&Documents)
	return Documents
}
