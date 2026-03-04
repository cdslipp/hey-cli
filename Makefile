BIN_DIR := ./bin
BINARY  := $(BIN_DIR)/hey
MODULE  := ./cmd/hey

.PHONY: build test lint clean install

build:
	@mkdir -p $(BIN_DIR)
	go build -o $(BINARY) $(MODULE)

test:
	go test ./...

lint:
	golangci-lint run ./...

clean:
	rm -rf $(BIN_DIR)

INSTALL_DIR := /usr/local/bin

install: build
	sudo install $(BINARY) $(INSTALL_DIR)/hey
