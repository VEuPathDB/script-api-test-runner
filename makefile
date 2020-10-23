VERSION=$(shell git describe --tags)

build:
	@env CGO_ENABLED=0 GOOS=linux go build -o bin/run -ldflags "-X 'main.version=${VERSION}'" cmd/main.go

travis:
	@env CGO_ENABLED=0 GOOS=linux go build -o bin/run -ldflags "-X 'main.version=${VERSION}'" cmd/main.go
	cd bin && tar -czf wdk-api-test-runner.linux.${TRAVIS_TAG}.tar.gz run && rm run

	@env CGO_ENABLED=0 GOOS=darwin go build -o bin/run -ldflags "-X main.version=${VERSION}" cmd/main.go
	cd bin && tar -czf wdk-api-test-runner.darwin.${TRAVIS_TAG}.tar.gz run && rm run

	@env CGO_ENABLED=0 GOOS=windows go build -o bin/run -ldflags "-X main.version=${VERSION}" cmd/main.go
	cd bin && tar -czf wdk-api-test-runner.windows.${TRAVIS_TAG}.tar.gz run && rm run
