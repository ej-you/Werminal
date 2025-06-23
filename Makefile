source = ./cmd/app/main.go


dev: swagger
	go run $(source)

lint:
	golangci-lint run -c ./.golangci.yml ./...
