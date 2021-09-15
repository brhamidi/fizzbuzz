package main

import (
	"github.com/brhamidi/fizzbuzz/pkg/http"
	"github.com/brhamidi/fizzbuzz/pkg/logger"
	"github.com/brhamidi/fizzbuzz/pkg/storage"
)

const appName = "fizzbuzz"

func main() {
	log := logger.NewLogger(appName)

	c, err := config()
	if err != nil {
		log.Fatal(err)
	}

	s := storage.NewInmemory()

	server := http.NewServer(c.Env, s, log)
	server.Run(":" + c.Port)
}
