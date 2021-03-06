.PHONY: docs deploy

docs:
	python3 ./docs/aws_diagram.py

deploy:
	# deploy server
	rm -rf ./bin
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w  -X github.com/drakejin/fiber-aws-serverless/config/config.Release=$(git log --format='%H' -n 1)" -o bin/lambda lambda/main.go
	SLS_DEBUG=* sls deploy --verbose
