package repo

import (
	"Library_project/other"
	"fmt"
	"strconv"
)

type Author struct {
	AuthorId    uint16 `json:"author_id"`
	AuthorName  string `json:"author_name" valid:"required"`
	AuthorImage string `json:"author_image" valid:"required"`
}

func GetAuthorsFromDB(Authors *[]Author, page string, limit string) {
	db := other.ConnectDB()
	defer db.Close()
	p, _ := strconv.Atoi(page)
	pageForSql := (p - 1) * 5
	l, _ := strconv.Atoi(limit)
	get, err := db.Query(fmt.Sprintf("Select * from `authors` LIMIT %d OFFSET %d", l, pageForSql))
	other.CheckErr(err)

	for get.Next() {
		var author Author
		err = get.Scan(&author.AuthorId, &author.AuthorName, &author.AuthorImage)
		other.CheckErr(err)
		*Authors = append(*Authors, author)
	}
}

func SaveAuthorInDB(author Author) {
	db := other.ConnectDB()
	defer db.Close()

	ins, err := db.Query(fmt.Sprintf("INSERT INTO `authors` (`author_name`,`author_image`) VALUES ('%s','%s')", author.AuthorName, author.AuthorImage))
	other.CheckErr(err)
	defer ins.Close()
}
