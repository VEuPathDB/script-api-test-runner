VERSION=$(shell git describe --tags)

build:
	env CGO_ENABLED=0 GOOS=linux go build -o bin/run -ldflags "-X internal.x.Version=${VERSION}" cmd/main.go
