package repo

import (
	"Library_project/other"
	"fmt"
)

type Author struct {
	AuthorId    string `json:"author_id"`
	AuthorName  string `json:"author_name"`
	AuthorImage string `json:"author_image"`
}

func GetAuthorsFromDB(Authors *[]Author) {
	db := other.ConnectDB()
	defer db.Close()

	get, err := db.Query("Select * from `authors`")
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
