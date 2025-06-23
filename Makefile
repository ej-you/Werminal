source = ./cmd/app/main.go


dev:
	go run $(source)

lint:
	golangci-lint run -c ./.golangci.yml ./...
