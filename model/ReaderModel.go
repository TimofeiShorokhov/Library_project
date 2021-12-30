package model

import (
	"Library_project/repo"
	"log"
)

type Reader struct {
	Id        uint16 `json:"id"`
	Name      string `json:"name"`
	Birthdate string `json:"birthdate"`
	Adress    string `json:"adress"`
	Surname   string `json:"surname"`
	Email     string `json:"email"`
	Debt      uint16 `json:"debt"`
}

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
