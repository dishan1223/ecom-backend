BINARY_NAME=ecom
CMD_DIR=./cmd/api

.PHONY: build
build:
	go tool templ generate
	go build -o $(BINARY_NAME) $(CMD_DIR)

.PHONY: run
run:
	go run $(CMD_DIR)/main.go

render:
	go tool templ generate

start:
	@echo "Starting..."
	go tool templ generate
	go run $(CMD_DIR)/main.go
