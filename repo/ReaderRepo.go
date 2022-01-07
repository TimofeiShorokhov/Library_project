package repo

import (
	"Library_project/other"
	"database/sql"
	"fmt"
	"strconv"
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

func IncreaseReaderDebtInDb(surname string, value uint16) {
	db := other.ConnectDB()
	defer db.Close()
	updDebt := db.QueryRow("UPDATE `readers` set debt = debt+? WHERE surname = ?", value, surname)
	updDebt.Err()
}

func DecreaseReaderDebtInDb(surname string) {
	db := other.ConnectDB()
	defer db.Close()
	updDebt := db.QueryRow("UPDATE `readers` set debt = debt-1 WHERE surname = ?", surname)
	updDebt.Err()
}

func SearchReaderInDb(surname string) bool {
	db := other.ConnectDB()
	defer db.Close()
	var reader Reader
	resReader := db.QueryRow("Select * from `readers` WHERE surname = ?", surname)
	errReader := resReader.Scan(&reader.Id, &reader.Name, &reader.Birthdate, &reader.Adress, &reader.Surname, &reader.Email, &reader.Debt)
	if errReader == sql.ErrNoRows {
		return false
	}
	return true
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

func GetReaderFromDBWithPages(Readers *[]Reader, page string) {
	db := other.ConnectDB()
	defer db.Close()
	p, _ := strconv.Atoi(page)
	get, err := db.Query(fmt.Sprintf("Select * from `readers` order by name LIMIT 5 OFFSET %d", (p-1)*5))
	other.CheckErr(err)

	for get.Next() {
		var reader Reader
		err = get.Scan(&reader.Id, &reader.Name, &reader.Birthdate, &reader.Adress, &reader.Surname, &reader.Email, &reader.Debt)
		other.CheckErr(err)
		*Readers = append(*Readers, reader)
	}
}
