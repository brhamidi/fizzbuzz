package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/kelseyhightower/envconfig"
)

type conf struct {
	// App config
	Port string `required:"true"`
	Env  string `required:"true" validate:"eq=debug|eq=release"`
	// Postgres config
	PGUser     string `required:"true" split_words:"true"`
	PGName     string `required:"true" split_words:"true"`
	PGPassword string `required:"true" split_words:"true"`
	PGHost     string `required:"true" split_words:"true"`
}

func config() (conf, error) {
	var c conf

	if err := envconfig.Process("", &c); err != nil {
		return conf{}, err
	}

	if err := validator.New().Struct(&c); err != nil {
		return conf{}, err
	}

	return c, nil
}
