.PHONY: build test

default: build

build:
	go build .

test:
	ginkgo -r .
