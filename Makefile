.PHONY: build test help default

default: build

help:
	@echo 'Management commands for qos-data-store:'
	@echo
	@echo 'Usage:'
	@echo '    make build           Compile the project.'
	@echo '    make get-deps        runs glide install, mostly used for ci.'
	@echo '    make build-docker    Build docker image'
	@echo '    make push            Push tagged images to registry'
	@echo '    make test            Run tests'
	@echo '    make dev             Run the program.'
	@echo

build:
	CGO_ENABLED=0 go build -a -installsuffix cgo

get-deps:
	glide install

build-docker:
	sudo docker build . -t hyperpilot/qos-data-store

push:
	sudo docker push hyperpilot/qos-data-store:latest

test:
	go test $(glide nv)

dev: build
	./qos-data-store --config documents/dev.config
