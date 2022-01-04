package model

import (
	"Library_project/repo"
	"log"
)

func SaveReader(reader *repo.Reader) {

	if reader.Name == "" || reader.Surname == "" || reader.Email == "" || reader.Birthdate == "" {
		log.Println("Не все поля заполнены")
	} else {
		repo.SaveReaderInDB(*reader)
	}
}

func GetReaders(Readers []repo.Reader) []repo.Reader {
	Readers = []repo.Reader{}
	repo.GetReaderFromDB(&Readers)
	return Readers
}
