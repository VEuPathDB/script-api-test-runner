VERSION=$(shell git describe --tags)

build:
	env CGO_ENABLED=0 GOOS=linux go build -o bin/run -ldflags "-X 'main.version=${VERSION}'" cmd/main.go
