package tmdb

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPopularMoviesAsJson(t *testing.T) {
	request := PopularMovieRequest{
		ApiKey:   "abcdefg",
		Language: "en",
		Page:     1,
	}

	bytes, err := json.Marshal(request)
	assert.Nil(t, err)
	assert.NotNil(t, bytes)

	fmt.Println(string(bytes))
}
