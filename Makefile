TARGET=ganzi
LDFLAGS=-ldflags "-X=main.Build=${BUILD}"
SRCS=main.go
BUILD_OPT=-race
ifndef $(GOPATH)
    GOPATH=$(shell go env GOPATH)
    export GOPATH
endif

.PHONY:  run clean build build-linux

run: build
	@./build/${TARGET}

build: check
	@mkdir -p build
	@go build ${BUILD_OPT} ${LDFLAGS} -o build/${TARGET} ${SRCS}

build-linux: check
	@mkdir -p build
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build ${LDFLAGS} -o build/${TARGET}-linux-amd64 ${SRCS}

check:
	@go vet ./...
	@go fmt ./...

test:
	go test ./...

clean:
	rm -fr build/${TARGET}*