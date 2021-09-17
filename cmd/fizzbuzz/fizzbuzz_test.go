package main

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {
	t.Run("should return an errParseEnv error", func(t *testing.T) {
		_, err := config()

		assert.Equal(t, true, errors.Is(err, errParseEnv))
	})

	t.Run("should return an errParseEnv error because validator failed", func(t *testing.T) {
		t.Setenv("PORT", "v")
		t.Setenv("ENV", "v")
		t.Setenv("REDIS_HOST", "v")
		t.Setenv("STORE_MODE", "v")
		t.Setenv("REDIS_PORT", "v")

		_, err := config()
		assert.Equal(t, true, errors.Is(err, errParseEnv))
	})

	t.Run("should be ok", func(t *testing.T) {
		t.Setenv("PORT", "v")
		t.Setenv("ENV", "debug")
		t.Setenv("STORE_MODE", "persistant")
		t.Setenv("REDIS_HOST", "v")
		t.Setenv("REDIS_PORT", "v")

		c, err := config()

		assert.NoError(t, err)
		assert.Equal(t, conf{"v", "debug", "persistant", "v", "v"}, c)
	})
}
