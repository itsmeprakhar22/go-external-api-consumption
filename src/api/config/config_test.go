package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	testAPITmdbAccessToken = "SECRET_ACCESS_TOKEN"
)

func TestConstants(t *testing.T) {
	assert.EqualValues(t, "SECRET_ACCESS_TOKEN", tmdbAccessToken)
}
func TestGetTmdbAccessToken(t *testing.T) {
	assert.EqualValues(t, os.Getenv(testAPITmdbAccessToken), tmdbAccessToken)
}
