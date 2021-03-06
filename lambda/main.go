package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	adapter "github.com/awslabs/aws-lambda-go-api-proxy/fiber"
	"github.com/rs/zerolog/log"

	"github.com/drakejin/fiber-aws-serverless/config"
	"github.com/drakejin/fiber-aws-serverless/internal/app"
	"github.com/drakejin/fiber-aws-serverless/internal/container"
)

func main() {
	cfg, err := config.New(true)
	if err != nil {
		log.Fatal().Err(err).Send()
	}

	cont, err := container.New(cfg)
	if err != nil {
		log.Fatal().Err(err).Send()
	}

	appHttp := app.NewHTTP(cont)
	lambdaApp := adapter.New(appHttp)
	lambda.Start(lambdaApp.ProxyWithContext)
}
