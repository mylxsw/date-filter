
build:
	go build -o ./bin/date-filter

build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./bin/date-filter-linux

clean:
	rm -fr ./bin/date-filter-linux ./bin/date-filter

include Makefile.local
