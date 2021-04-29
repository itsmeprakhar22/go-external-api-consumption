package provider

import (
	"GolangWorkspace/go-consuming-apis/src/api/clients/rest_clients"
	"GolangWorkspace/go-consuming-apis/src/api/domain/tmdb"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	headerAuthorization       = "Authorization"
	headerAuthorizationFormat = "Bearer %s"

	urlTmdbBase string = "http://api.themoviedb.org/3/movie/popular"
)

func getAuthorizationHeader(accessToken string) string {
	return fmt.Sprintf(headerAuthorizationFormat, accessToken)
}

func GetPopularMovies(accessToken string, request tmdb.PopularMovieRequest) (*tmdb.PopularMovieResponse, *tmdb.PopularMovieError) {

	headers := http.Header{}
	headers.Set(headerAuthorization, getAuthorizationHeader(accessToken))
	params := make(map[string]string)
	params["api_key"] = request.ApiKey
	params["language"] = request.Language
	params["page"] = fmt.Sprint(request.Page)

	response, err := rest_clients.Get(urlTmdbBase, params, headers)

	if err != nil {
		// log.Println("Error while fetching movie data")
		return nil, &tmdb.PopularMovieError{
			StatusCode:    http.StatusInternalServerError,
			StatusMessage: err.Error(),
			Success:       false,
		}
	}

	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		// log.Println("Invalid response body")
		return nil, &tmdb.PopularMovieError{
			StatusCode:    http.StatusInternalServerError,
			StatusMessage: "Invalid response body",
			Success:       false,
		}
	}

	defer response.Body.Close()

	if response.StatusCode > 299 {
		var errorResponse tmdb.PopularMovieError
		if err := json.Unmarshal(bytes, &errorResponse); err != nil {
			return nil, &tmdb.PopularMovieError{
				StatusCode:    http.StatusInternalServerError,
				StatusMessage: "Invalid json response body",
				Success:       false,
			}
		}
		errorResponse.StatusCode = response.StatusCode
		return nil, &errorResponse
	}

	var result tmdb.PopularMovieResponse
	if err = json.Unmarshal(bytes, &result); err != nil {
		return nil, &tmdb.PopularMovieError{
			StatusCode:    http.StatusInternalServerError,
			StatusMessage: "Error in unmarshal response body",
			Success:       false,
		}
	}
	return &result, nil
}
