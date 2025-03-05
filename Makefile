BIN_NAME := commitsmith

.PHONY: all build test clean
all: build

build:
	go build -o $(BIN_NAME)

test:
	go test .

clean:
	@rm -rf $(BIN_NAME)