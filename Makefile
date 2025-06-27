# + ------ +
# + server +
# + ------ +

source = ./cmd/app/main.go

dev:
	go run $(source)

lint:
	golangci-lint run -c ./.golangci.yml ./...

# + ------ +
# + client +
# + ------ +

client:
	cd ./website && npx http-server -p 3000

ts:
	cd ./website && tsc
