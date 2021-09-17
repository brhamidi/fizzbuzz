package main

import (
	"errors"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TODO update to go1.17 to use the new function testing.SetEnv

func TestConfig(t *testing.T) {
	t.Run("should return an errParseEnv error", func(t *testing.T) {
		_, err := config()

		assert.Equal(t, true, errors.Is(err, errParseEnv))
	})

	t.Run("should return an errParseEnv error because validator failed", func(t *testing.T) {
		os.Setenv("PORT", "v")
		os.Setenv("ENV", "v")
		os.Setenv("REDIS_HOST", "v")
		os.Setenv("STORE_MODE", "v")
		os.Setenv("REDIS_PORT", "v")
		defer os.Unsetenv("PORT")
		defer os.Unsetenv("ENV")
		defer os.Unsetenv("STORE_MODE")
		defer os.Unsetenv("REDIS_HOST")
		defer os.Unsetenv("REDIS_PORT")

		_, err := config()
		assert.Equal(t, true, errors.Is(err, errParseEnv))
	})

	t.Run("should be ok", func(t *testing.T) {
		os.Setenv("PORT", "v")
		os.Setenv("ENV", "debug")
		os.Setenv("STORE_MODE", "persistant")
		os.Setenv("REDIS_HOST", "v")
		os.Setenv("REDIS_PORT", "v")
		defer os.Unsetenv("PORT")
		defer os.Unsetenv("ENV")
		defer os.Unsetenv("REDIS_HOST")
		defer os.Unsetenv("REDIS_PORT")

		c, err := config()

		assert.NoError(t, err)
		assert.Equal(t, conf{"v", "debug", "persistant", "v", "v"}, c)
	})
}
