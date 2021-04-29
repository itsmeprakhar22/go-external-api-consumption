package app

import (
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApp() {
	mapUrls()
	if err := router.Run(":8000"); err != nil {
		panic(err)
	}
}
