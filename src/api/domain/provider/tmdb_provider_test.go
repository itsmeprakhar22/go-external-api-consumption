package provider

import (
	"GolangWorkspace/go-consuming-apis/src/api/domain/tmdb"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAuthorizationHeader(t *testing.T) {
	header := getAuthorizationHeader("abc123")
	assert.EqualValues(t, "Bearer abc123", header)
}

func TestGetPopularMovies(t *testing.T) {
	r := tmdb.PopularMovieRequest{
		ApiKey:   "abcdefg",
		Language: "en",
		Page:     1,
	}
	res, err := GetPopularMovies("YOUR_ACCESS_TOKEN", r)
	if err != nil {
		log.Print("Error occured!")
		log.Println(err)
	}
	assert.NotNil(t, res)
}
