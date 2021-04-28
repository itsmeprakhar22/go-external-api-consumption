package services

import "GolangWorkspace/go-consuming-apis/src/api/domain/tmdb"

type tmdbService struct {
}

type tmdbServiceInterface interface {
}

var (
	TmdbService tmdbServiceInterface
)

func init() {
	TmdbService = &tmdbService{}
}

func (s *tmdbService) GetPopularMovies(request tmdb.PopularMovieRequest) {

}
