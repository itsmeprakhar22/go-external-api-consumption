package tmdb

import (
	"GolangWorkspace/go-consuming-apis/src/api/domain/provider"
	"testing"
)

// func TestGetPopularMoviesAsJson(t *testing.T) {
// 	request := PopularMovieRequest{
// 		ApiKey:   "abcdefg",
// 		Language: "en",
// 		Page:     1,
// 	}

// 	bytes, err := json.Marshal(request)
// 	assert.Nil(t, err)
// 	assert.NotNil(t, bytes)

// 	fmt.Println(string(bytes))
// }

func TestGetPopularMovies(t *testing.T) {
	r := PopularMovieRequest{
		ApiKey:   "abcdefg",
		Language: "en",
		Page:     1,
	}
	provider.GetPopularMovies("", r)
}
