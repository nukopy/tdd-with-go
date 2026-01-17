SOURCES ?= $(shell find . -type f \( -name "*.go" -o -name "go.mod" -o -name "go.sum" \) -print)

.DEFAULT_GOAL := help

# ------------------------------------------------------------
# Setup
# ------------------------------------------------------------

# .PHONY: init
# init: ## Download and install go mod dependencies
# 	go mod download
# 	# go install github.com/google/wire/cmd/wire@v0.7.0
# 	# go install github.com/golang/mock/mockgen@v1.6.0

# ------------------------------------------------------------
# Build
# ------------------------------------------------------------

tdd-with-go: $(SOURCES) ## Build tdd-with-go binary
	CGO_ENABLED=0 go build -o tdd-with-go -ldflags "-s -w" -trimpath

# ------------------------------------------------------------
# CI
# ------------------------------------------------------------

.PHONY: fmt
fmt: ## Format go files
	golangci-lint fmt ./...

.PHONY: golangci-lint
golangci-lint: ## Lint and check format go files
	@golangci-lint run

.PHONY: lint-and-fmt
lint-and-fmt: ## Lint and format go files
	-@make golangci-lint

.PHONY: lint
lint: ## (alias: lint-and-fmt) Lint and format go files
	-@make lint-and-fmt

.PHONY: test
test: ## Run test
	go test ./... -race -shuffle=on -vet=off

.PHONY: test-with-coverage
test-with-coverage: ## Run test with coverage
	go test ./... -coverprofile=coverage.txt -race -shuffle=on -vet=off

# ------------------------------------------------------------
# Docker
# ------------------------------------------------------------

# .PHONY: up
# up: ## Build and start the app containers
# 	@docker compose up -d --build

# .PHONY: down
# down: ## Stop and remove app containers
# 	@docker compose down

# ------------------------------------------------------------
# Codegen
# ------------------------------------------------------------

# .PHONY: gogen
# gogen: ## Generate auto-generated go files
# 	go generate ./...

# ------------------------------------------------------------
# Help
# ------------------------------------------------------------

.PHONY: help
help: ## Display this help screen
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'
