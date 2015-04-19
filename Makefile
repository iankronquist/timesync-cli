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
	go test -coverprofile=c.out ./timesync
	go tool cover -html=c.out

clean:
	rm -rf build
	rm -f timesync-cli
