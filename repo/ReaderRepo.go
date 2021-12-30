package repo

import (
	"Library_project/other"
	"fmt"
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

func SaveReaderInDB(reader Reader) {
	db := other.ConnectDB()
	defer db.Close()

	ins, err := db.Query(fmt.Sprintf("INSERT INTO `readers` (`name`,`surname`,`birthdate`,`email`,`adress`) VALUES ('%s','%s','%s','%s','%s')", reader.Name, reader.Surname, reader.Birthdate, reader.Email, reader.Adress))
	other.CheckErr(err)
	defer ins.Close()
}

func GetReaderFromDB(Readers *[]Reader) {
	db := other.ConnectDB()
	defer db.Close()

	get, err := db.Query("Select * from `readers` order by name")
	other.CheckErr(err)

	for get.Next() {
		var reader Reader
		err = get.Scan(&reader.Id, &reader.Name, &reader.Birthdate, &reader.Adress, &reader.Surname, &reader.Email, &reader.Debt)
		other.CheckErr(err)
		*Readers = append(*Readers, reader)
	}
}
