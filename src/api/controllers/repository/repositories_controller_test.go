package repository

import (
	"GolangWorkspace/go-consuming-apis/src/api/clients/rest_clients"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	rest_clients.StartMockups()
	os.Exit(m.Run())
}

func TestGetPopularMoviesInvalidJson(t *testing.T) {
	response := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(response)

	request, _ := http.NewRequest(http.MethodGet, "http://api.themoviedb.org/3/movie/popular", nil)
	c.Request = request
	rest_clients.AddMockup(rest_clients.Mock{
		Url:        "http://api.themoviedb.org/3/movie/popular",
		HttpMethod: http.MethodGet,
		Error:      errors.New("invalid json response"),
	})
	GetPopularMovies(c)
	assert.EqualValues(t, http.StatusInternalServerError, response.Code)
	fmt.Println(response.Body.String())
}
