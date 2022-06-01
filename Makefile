OUTPUT_PATH ?= ./go-gin-client

build:
	CGO_ENABLED=0 go build -o $(OUTPUT_PATH)

run:
	go run .
