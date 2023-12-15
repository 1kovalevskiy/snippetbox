package main

import (
	"log"

	"github.com/1kovalevskiy/snippetbox/config"
	"github.com/1kovalevskiy/snippetbox/internal/app"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	// Run
	app.Run(cfg)

}
