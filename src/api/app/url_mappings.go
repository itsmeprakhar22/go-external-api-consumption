package app

import (
	"GolangWorkspace/go-consuming-apis/src/api/controllers/repository"
)

func mapUrls() {
	router.GET("/movies", repository.GetPopularMovies)
}
