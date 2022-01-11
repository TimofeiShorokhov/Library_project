package model

import (
	"Library_project/other"
	"Library_project/repo"
	"github.com/asaskevich/govalidator"
)

func SaveReader(reader *repo.Reader) {
	result, err := govalidator.ValidateStruct(reader)
	other.CheckErr(err)
	if result == true {
		repo.SaveReaderInDB(*reader)
	}
}

func GetReaders(Readers []repo.Reader) []repo.Reader {
	Readers = []repo.Reader{}
	repo.GetReaderFromDB(&Readers)
	return Readers
}

func GetReadersWithPage(Readers []repo.Reader, page string) []repo.Reader {
	Readers = []repo.Reader{}
	repo.GetReaderFromDBWithPages(&Readers, page)
	return Readers
}
