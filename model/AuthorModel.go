package model

import (
	"Library_project/other"
	"Library_project/repo"
	"fmt"
	"log"
)

func GetAuthors(Authors []repo.Author) []repo.Author {
	Authors = []repo.Author{}
	repo.GetAuthorsFromDB(&Authors)
	return Authors
}

func SaveAuthor(author *repo.Author) {

	if author.AuthorName == "" {
		log.Println("Не все поля заполнены")
	} else {
		image := author.AuthorImage
		filepath := fmt.Sprintf("./images/author_img/%s.jpg", author.AuthorName)
		other.DownloadFile(filepath, image)
		author.AuthorImage = filepath
		repo.SaveAuthorInDB(*author)
	}

}
