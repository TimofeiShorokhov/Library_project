package repo

import "Library_project/other"

type Genre struct {
	GenreId   string `json:"genre_id"`
	GenreName string `json:"book_genre"`
}

func GetGenresFromDB(Genres *[]Genre) {
	db := other.ConnectDB()
	defer db.Close()

	get, err := db.Query("Select * from `genres`")
	other.CheckErr(err)

	for get.Next() {
		var genre Genre
		err = get.Scan(&genre.GenreId, &genre.GenreName)
		other.CheckErr(err)
		*Genres = append(*Genres, genre)
	}
}
