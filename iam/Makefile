SRCS = $(shell git ls-files '*.go' | grep -v '^vendor/')

BIN_NAME := zziam

# Default build target
GOOS := $(shell go env GOOS)
GOARCH := $(shell go env GOARCH)

.PHONY: run
#? run: Run the application
run:
	@go run cmd/$(BIN_NAME)/$(BIN_NAME).go

.PHONY: lint
#? lint: Run golangci-lint
lint:
	golangci-lint run

.PHONY: fmt
#? fmt: Format the Code
fmt:
	gofmt -s -l -w $(SRCS)

.PHONY: help
#? help: Get more info on make commands
help: Makefile
	@echo " Choose a command run:"
	@sed -n 's/^#?//p' $< | column -t -s ':' |  sort | sed -e 's/^/ /'