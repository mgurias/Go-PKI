.PHONY: clean test security build run

APP_NAME = apiserver
BUILD_DIR = ${PWD}/build

clean:
	rm -rf ./build

security:
	gosec -quiet ./...

test: security
	go test -v -timeout 30s -coverprofile=cover.out -cover ./...
	go tool cover -func=cover.out

build: clean test
	CGO_ENABLED=0 go build -ldflags="-w -s" -o ${BUILD_DIR}/${APP_NAME} main.go

run: ${GOPATH}/bin/swag build
	${BUILD_DIR}/${APP_NAME}

swag:
	${GOPATH}/bin/swag init
