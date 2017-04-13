.PHONY: all build test

all: test

build:
	go get .

test: build
	./test.sh
