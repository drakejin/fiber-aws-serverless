package main

import (
	"github.com/rs/zerolog/log"

	"github.com/drakejin/fiber-aws-serverless/cmd"
	"github.com/drakejin/fiber-aws-serverless/config"
)

func main() {
	var isServerless = false
	cfg, err := config.New(isServerless)
	if err != nil {
		log.Fatal().Err(err).Send()
	}
	cmd.Start(cfg)
}
