package services

import (
	"GolangWorkspace/go-consuming-apis/src/api/clients/rest_clients"
	"GolangWorkspace/go-consuming-apis/src/api/domain/repositories"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	rest_clients.StartMockups()
	os.Exit(m.Run())
}

func TestGetPopularMoviesErrorFromTMDB(t *testing.T) {
	rest_clients.FlushMockups()
	invalidCloser, _ := os.Open("-asf3")
	rest_clients.AddMockup(rest_clients.Mock{
		Url:        "http://api.themoviedb.org/3/movie/popular",
		HttpMethod: http.MethodGet,
		Response: &http.Response{
			StatusCode: http.StatusCreated,
			Body:       invalidCloser,
		},
	})
}

func TestGetPopularMoviesNoError(t *testing.T) {
	rest_clients.FlushMockups()

	rest_clients.AddMockup(rest_clients.Mock{
		Url:        "http://api.themoviedb.org/3/movie/popular",
		HttpMethod: http.MethodGet,
		Response: &http.Response{
			StatusCode: http.StatusAccepted,
			Body:       ioutil.NopCloser(strings.NewReader(`{"page":2,"results":[{"backdrop_path":"/9yBVqNruk6Ykrwc32qrK2TIE5xw.jpg","overview":"Washed-up USA fighter Cole Young, unaware of his heritage, and hunted by Emperor Shang Tsung's best warrior, Sub-Zero, seeks out and trains with Earth's greatest champions as he prepares to stand against the enemies of Outworld in a high stakes battle for the universe.","release_date":"2021-04-07","title":"Mortal Kombat"}]}`)),
		},
	})
	request := repositories.GetPopularMoviesRequest{}
	request.Language = "ar"
	res, err := TmdbService.GetPopularMovies(request)

	assert.NotNil(t, res)
	assert.Nil(t, err)
	assert.EqualValues(t, "/9yBVqNruk6Ykrwc32qrK2TIE5xw.jpg", res.Result[0].BackdropPath)
	assert.EqualValues(t, 2, res.Page)
}

func TestGetPopularMoviesUnauthorizated(t *testing.T) {
	rest_clients.FlushMockups()

	rest_clients.AddMockup(rest_clients.Mock{
		Url:        "http://api.themoviedb.org/3/movie/popular",
		HttpMethod: http.MethodGet,
		Response: &http.Response{
			StatusCode: http.StatusUnauthorized,
			Body:       ioutil.NopCloser(strings.NewReader(`{"status_code":7,"status_message":"Invalid API key: You must be granted a valid key.","success":false}`)),
		},
	})
	request := repositories.GetPopularMoviesRequest{}
	request.Language = "ar"
	res, err := TmdbService.GetPopularMovies(request)
	assert.NotNil(t, err)
	assert.Nil(t, res)
	assert.EqualValues(t, http.StatusUnauthorized, err.Status())
	assert.EqualValues(t, "Invalid API key: You must be granted a valid key.", err.Message())
}
