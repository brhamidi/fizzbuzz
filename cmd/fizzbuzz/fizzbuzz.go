package main

import (
	"github.com/brhamidi/fizzbuzz/pkg/http"
	"github.com/brhamidi/fizzbuzz/pkg/logger"
	"github.com/brhamidi/fizzbuzz/pkg/storage"
)

func main() {
	log := logger.NewLogger("fizzbuzz")

	c, err := config()
	if err != nil {
		log.Fatal(err)
	}

	s, err := storage.NewPersistant()
	if err != nil {
		log.Fatal(err)
	}

	server := http.NewServer(c.Env, s, log)
	server.Run(":" + c.Port)
}
