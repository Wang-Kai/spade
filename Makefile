.PHONY: test go/test

build:
	@mkdir -p bin
	go build -v -o bin/spade_macos main.go
	GOOS=linux go build -v -o bin/spade_linux main.go


test: build
	bin/spade -t example/template -v example/values.yaml

go/test:
	go test -v ./...