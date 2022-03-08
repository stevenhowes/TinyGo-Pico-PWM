.DEFAULT_GOAL := install

fmt:
	go mod tidy
	go fmt ./...
.PHONY:fmt

lint: fmt
	golint ./...
.PHONY:lint

vet: fmt
# Need to work this out for tinygo..
#	go vet ./...
.PHONY:vet

install: vet
	tinygo flash -target=pico
.PHONY:install

