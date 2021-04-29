package config

import "os"

const (
	apiTmdbAccessToken = "SECRET_ACCESS_TOKEN"
)

var (
	tmdbAccessToken = os.Getenv(apiTmdbAccessToken)
)

func GetTmdbAccessToken() string {
	return tmdbAccessToken
}
