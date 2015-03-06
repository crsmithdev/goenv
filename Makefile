.PHONY: build test

default: build

deps:
	go get github.com/onsi/ginkgo/ginkgo
	go get github.com/stretchr/testify/assert

build:
	go build .

test:
	ginkgo -r .

test-auto: test
	ginkgo watch -r -cover .

test-cov:
	ginkgo -r -cover .
	echo "mode: atomic" > coverage.out
	@for file in $$(find . -name "*.coverprofile" ! -name "coverage.out"); do \
		cat $$file | grep -v "mode: atomic" | sed 's|^_.*'$$(pwd)'|.|g' >> coverage.out ; \
		rm $$file ; \
	done

html-cov:
	go tool cover -html=coverage.out
