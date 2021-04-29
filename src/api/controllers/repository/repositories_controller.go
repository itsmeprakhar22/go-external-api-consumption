package repository

import (
	"GolangWorkspace/go-consuming-apis/src/api/config"
	"GolangWorkspace/go-consuming-apis/src/api/domain/repositories"
	"GolangWorkspace/go-consuming-apis/src/api/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetPopularMovies(c *gin.Context) {
	var request repositories.GetPopularMoviesRequest
	var e error

	request.Language = c.Request.URL.Query().Get("language")
	request.ApiKey = config.GetTmdbAccessToken()
	request.Page, e = strconv.Atoi(c.Request.URL.Query().Get("page"))
	if e != nil {
		request.Page = 1
	}
	// request.Page = strconv.Atoi(c.Request.URL.Query().Get("page"))
	// if err := c.Request.URL.Query(); err != nil {
	// 	apiErr := utils.NewBadRequest("invalid json body")
	// 	c.JSON(apiErr.Status(), apiErr)
	// 	return
	// }
	result, err := services.TmdbService.GetPopularMovies(request)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusCreated, &result)
	return
}
