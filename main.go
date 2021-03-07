package main

import (
	"os"

	"github.com/rs/zerolog/log"

	"github.com/drakejin/fiber-aws-serverless/cmd"
	"github.com/drakejin/fiber-aws-serverless/config"
)

func main() {
	var useDotenv = true
	if os.Getenv("IS_DOCKER") == "true" {
		useDotenv = false
	}

	cfg, err := config.New(useDotenv)
	if err != nil {
		log.Fatal().Err(err).Send()
	}
	cmd.Start(cfg)
}
