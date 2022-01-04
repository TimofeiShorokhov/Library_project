package model

import "Library_project/repo"

func GetGenres(Genres []repo.Genre) []repo.Genre {
	Genres = []repo.Genre{}
	repo.GetGenresFromDB(&Genres)
	return Genres
}
