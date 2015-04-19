all: build

setup:
	mkdir -p ${GOPATH}/src/${ORGPATH}
	mkdir -p build

build: setup
	go build timesync-cli.go

run:
	go run timesync-cli.go

test:
	go test ./timesync

install-deps:
	echo install dependencies...

cover:
	go test -coverprofile=coverage.out ./timesync
	go tool cover -html=coverage.out

clean:
	rm -rf build
	rm -f timesync-cli
	rm -f coverage.out
