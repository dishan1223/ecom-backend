BINARY_NAME=ecom
CMD_DIR=./cmd/api

.PHONY: build
build:
	go build -o $(BINARY_NAME) $(CMD_DIR)

.PHONY: run
run:
	@echo "Running $(BINARY_NAME)..."
	go run $(CMD_DIR)/main.go
