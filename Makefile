.PHONY: all lint test build
.SILENT:

lint:
	golangci-lint run

swagger:
	swag init -g ./internal/app/idler-service/app.go

clean:
	go clean -modcache

build: test
	go build -o ./.bin/app ./cmd/idler-service/main.go

test:
	go test -v ./test/...