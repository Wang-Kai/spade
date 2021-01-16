.PHONY: test go/test

build:
	@mkdir -p bin
	@go build -v -o bin/spade main.go

test: build
	bin/spade -t example/template -v example/values.yaml

go/test:
	go test -v ./...