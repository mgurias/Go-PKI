.PHONY: clean test security build run

APP_NAME = apiserver
BUILD_DIR = ${PWD}/build

clean:
	rm -rf ./build

build: clean
	   CGO_ENABLED=0 go build -ldflags="-w -s" -o ${BUILD_DIR}/${APP_NAME} main.go

run: ${GOPATH}/bin/swag build ${BUILD_DIR}/${APP_NAME}

swag:${GOPATH}/bin/swag init
