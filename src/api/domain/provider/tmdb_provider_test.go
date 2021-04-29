package provider

import (
	"GolangWorkspace/go-consuming-apis/src/api/clients/rest_clients"
	"GolangWorkspace/go-consuming-apis/src/api/domain/tmdb"
	"errors"
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

func TestGetAuthorizationHeader(t *testing.T) {
	header := getAuthorizationHeader("abc123")
	assert.EqualValues(t, "Bearer abc123", header)
}

func TestGetPopularMovies(t *testing.T) {
	rest_clients.FlushMockups()
	rest_clients.AddMockup(rest_clients.Mock{
		Url:        "http://api.themoviedb.org/3/movie/popular",
		HttpMethod: http.MethodGet,
		Error:      errors.New("invalid json response"),
	})

	res, err := GetPopularMovies("YOUR_ACCESS_TOKEN", tmdb.PopularMovieRequest{})
	assert.NotNil(t, err)
	assert.Nil(t, res)
}

func TestGetPopularMoviesInvalidResponseBody(t *testing.T) {
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

	res, err := GetPopularMovies("YOUR_TOKEN", tmdb.PopularMovieRequest{})

	assert.NotNil(t, err)
	assert.Nil(t, res)
	assert.EqualValues(t, http.StatusInternalServerError, err.StatusCode)
	assert.EqualValues(t, "Invalid response body", err.StatusMessage)
}

func TestGetPopularMoviesInvalidSuccessResponse(t *testing.T) {
	rest_clients.FlushMockups()

	rest_clients.AddMockup(rest_clients.Mock{
		Url:        "http://api.themoviedb.org/3/movie/popular",
		HttpMethod: http.MethodGet,
		Response: &http.Response{
			StatusCode: http.StatusCreated,
			Body:       ioutil.NopCloser(strings.NewReader(`{"page": "123"}`)),
		},
	})

	res, err := GetPopularMovies("YOUR_TOKEN", tmdb.PopularMovieRequest{})

	assert.NotNil(t, err)
	assert.Nil(t, res)
	assert.EqualValues(t, http.StatusInternalServerError, err.StatusCode)
	assert.EqualValues(t, "Error in unmarshal response body", err.StatusMessage)
}

func TestGetPopularMoviesRequiresAuthentication(t *testing.T) {
	rest_clients.FlushMockups()

	rest_clients.AddMockup(rest_clients.Mock{
		Url:        "http://api.themoviedb.org/3/movie/popular",
		HttpMethod: http.MethodGet,
		Response: &http.Response{
			StatusCode: http.StatusUnauthorized,
			Body:       ioutil.NopCloser(strings.NewReader(`{"status_code":7,"status_message":"Invalid API key: You must be granted a valid key.","success":false}`)),
		},
	})

	res, err := GetPopularMovies("YOUR_TOKEN", tmdb.PopularMovieRequest{})

	assert.NotNil(t, err)
	assert.Nil(t, res)
	assert.EqualValues(t, http.StatusUnauthorized, err.StatusCode)
	assert.EqualValues(t, "Invalid API key: You must be granted a valid key.", err.StatusMessage)
}

func TestGetPopularMoviesInvalidErrorInterface(t *testing.T) {
	rest_clients.FlushMockups()

	rest_clients.AddMockup(rest_clients.Mock{
		Url:        "http://api.themoviedb.org/3/movie/popular",
		HttpMethod: http.MethodGet,
		Response: &http.Response{
			StatusCode: http.StatusUnauthorized,
			Body:       ioutil.NopCloser(strings.NewReader(`{"status_code":"1123",","success":false}`)),
		},
	})

	res, err := GetPopularMovies("YOUR_TOKEN", tmdb.PopularMovieRequest{})

	assert.NotNil(t, err)
	assert.Nil(t, res)
	assert.EqualValues(t, http.StatusInternalServerError, err.StatusCode)
	assert.EqualValues(t, "Invalid json response body", err.StatusMessage)
}

func TestGetPopularMoviesNoError(t *testing.T) {
	rest_clients.FlushMockups()

	rest_clients.AddMockup(rest_clients.Mock{
		Url:        "http://api.themoviedb.org/3/movie/popular",
		HttpMethod: http.MethodGet,
		Response: &http.Response{
			StatusCode: http.StatusAccepted,
			Body:       ioutil.NopCloser(strings.NewReader(`{"page":1,"results":[{"backdrop_path":"/9yBVqNruk6Ykrwc32qrK2TIE5xw.jpg","overview":"Washed-up MMA fighter Cole Young, unaware of his heritage, and hunted by Emperor Shang Tsung's best warrior, Sub-Zero, seeks out and trains with Earth's greatest champions as he prepares to stand against the enemies of Outworld in a high stakes battle for the universe.","release_date":"2021-04-07","title":"Mortal Kombat"}]}`)),
		},
	})

	res, err := GetPopularMovies("YOUR_TOKEN", tmdb.PopularMovieRequest{})

	assert.NotNil(t, res)
	assert.Nil(t, err)
	assert.EqualValues(t, "/9yBVqNruk6Ykrwc32qrK2TIE5xw.jpg", res.Result[0].BackdropPath)
	assert.EqualValues(t, 1, res.Page)
}
