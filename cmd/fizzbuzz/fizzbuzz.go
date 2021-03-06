package main

import (
	"errors"
	"fmt"

	"github.com/brhamidi/fizzbuzz/pkg/http"
	"github.com/brhamidi/fizzbuzz/pkg/logger"
	"github.com/brhamidi/fizzbuzz/pkg/storage"
	"github.com/go-playground/validator/v10"
	"github.com/kelseyhightower/envconfig"
)

const appName = "fizzbuzz"

var (
	errParseEnv      = errors.New("failed to parse environment variable")
	errStoreInstance = errors.New("failed to instanciate storage")
)

type conf struct {
	// App config
	Port      string `required:"true"`
	Env       string `required:"true" validate:"eq=debug|eq=release"`
	StoreMode string `required:"true" validate:"eq=inmemory|eq=persistant" split_words:"true"`
	// Redis config
	RedisHost string `required:"true" split_words:"true"`
	RedisPort string `required:"true" split_words:"true"`
}

func config() (conf, error) {
	var c conf

	if err := envconfig.Process("", &c); err != nil {
		return conf{}, fmt.Errorf("%w: %s", errParseEnv, err)
	}

	if err := validator.New().Struct(&c); err != nil {
		return conf{}, fmt.Errorf("%w: %s", errParseEnv, err)
	}

	return c, nil
}

func main() {
	log := logger.NewLogger(appName)

	c, err := config()
	if err != nil {
		log.Fatal(err)
	}

	var store storage.Storage

	if c.StoreMode == "inmemory" {
		store = storage.NewInmemory()
	} else {
		store, err = storage.NewPersistant(c.RedisHost, c.RedisPort)
		if err != nil {
			log.Fatal(fmt.Errorf("%w: %s", errStoreInstance, err))
		}
	}

	server := http.NewServer(c.Env, store, log)
	server.Run(":" + c.Port)
}
