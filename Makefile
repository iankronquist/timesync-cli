GOPATH=${PWD}/gopath

all: build

setup:
	mkdir -p ${GOPATH}/src/${ORGPATH}

build: setup
	go build main.go

run:
	go run main.go

test:
	go test

install-deps:
	echo install dependencies...
