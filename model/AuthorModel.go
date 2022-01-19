package model

import (
	"Library_project/other"
	"Library_project/repo"
	"fmt"
	"github.com/asaskevich/govalidator"
)

func GetAuthors(Authors []repo.Author, page string, limit string) []repo.Author {
	Authors = []repo.Author{}
	repo.GetAuthorsFromDB(&Authors, page, limit)
	return Authors
}

func SaveAuthor(author *repo.Author) {
	res, err := govalidator.ValidateStruct(author)
	if res != true {
		other.CheckErr(err)
	} else {
		image := author.AuthorImage
		filepath := fmt.Sprintf("./images/author_img/%s.jpg", author.AuthorName)
		other.DownloadFile(filepath, image)
		author.AuthorImage = filepath
		repo.SaveAuthorInDB(*author)
	}

}
