package services

import (
	"GolangWorkspace/go-consuming-apis/src/api/config"
	"GolangWorkspace/go-consuming-apis/src/api/domain/provider"
	"GolangWorkspace/go-consuming-apis/src/api/domain/repositories"
	"GolangWorkspace/go-consuming-apis/src/api/domain/tmdb"
	"GolangWorkspace/go-consuming-apis/src/api/utils"
	"log"
	"strings"
)

type tmdbService struct {
}

type tmdbServiceInterface interface {
	GetPopularMovies(request repositories.GetPopularMoviesRequest) (*repositories.GetPopularMoviesResponse, utils.ApiError)
}

var (
	TmdbService tmdbServiceInterface
)

func init() {
	TmdbService = &tmdbService{}
}

func (s *tmdbService) GetPopularMovies(input repositories.GetPopularMoviesRequest) (*repositories.GetPopularMoviesResponse, utils.ApiError) {
	input.ApiKey = strings.TrimSpace(input.ApiKey)
	// if input.ApiKey == "" {
	// 	return nil, utils.NewBadRequest("Invalid api key")
	// }

	input.Language = strings.TrimSpace(input.Language)
	if input.Language == "" {
		input.Language = "en-US"
	}

	request := tmdb.PopularMovieRequest{
		ApiKey:   input.ApiKey,
		Language: input.Language,
		Page:     input.Page,
	}

	log.Println(config.GetTmdbAccessToken())
	response, err := provider.GetPopularMovies(config.GetTmdbAccessToken(), request)
	if err != nil {
		return nil, utils.NewApiError(err.StatusCode, err.StatusMessage)
	}

	popularMovieList := PopularMovieProviderToService(response.Result)
	result := repositories.GetPopularMoviesResponse{
		Page:   response.Page,
		Result: popularMovieList,
	}

	return &result, nil
}

func PopularMovieProviderToService(bs []tmdb.PopularMovie) []repositories.GetPopularMovie {
	var acc []repositories.GetPopularMovie

	for _, b := range bs {
		newGetPopularMovie := repositories.GetPopularMovie{
			Title:        b.Title,
			Overview:     b.Overview,
			BackdropPath: b.BackdropPath,
			ReleaseDate:  b.ReleaseDate,
		} // pulled out for clarity
		acc = append(acc, newGetPopularMovie)
	}

	return acc
}
