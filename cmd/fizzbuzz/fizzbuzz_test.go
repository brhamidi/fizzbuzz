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
		os.Setenv("PG_HOST", "v")
		os.Setenv("PG_USER", "v")
		os.Setenv("PG_NAME", "v")
		os.Setenv("PG_PASSWORD", "v")
		defer os.Unsetenv("PORT")
		defer os.Unsetenv("ENV")
		defer os.Unsetenv("PG_HOST")
		defer os.Unsetenv("PG_USER")
		defer os.Unsetenv("PG_NAME")
		defer os.Unsetenv("PG_PASSWORD")

		_, err := config()
		assert.Equal(t, true, errors.Is(err, errParseEnv))
	})

	t.Run("should be ok", func(t *testing.T) {
		os.Setenv("PORT", "v")
		os.Setenv("ENV", "debug")
		os.Setenv("PG_HOST", "v")
		os.Setenv("PG_USER", "v")
		os.Setenv("PG_NAME", "v")
		os.Setenv("PG_PASSWORD", "v")
		defer os.Unsetenv("PORT")
		defer os.Unsetenv("ENV")
		defer os.Unsetenv("PG_HOST")
		defer os.Unsetenv("PG_USER")
		defer os.Unsetenv("PG_NAME")
		defer os.Unsetenv("PG_PASSWORD")

		c, err := config()

		assert.NoError(t, err)
		assert.Equal(t, conf{"v", "debug", "v", "v", "v", "v"}, c)
	})
}
