package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/kelseyhightower/envconfig"
)

type conf struct {
	Port string `required:"true"`
	Env  string `required:"true" validate:"eq=debug|eq=release"`
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
