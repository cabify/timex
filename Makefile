.PHONY: test help fmt report-coveralls benchmark lint

help: ## Show the help text
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "    \033[36m%-20s\033[93m %s\n", $$1, $$2}'

test: ## Run unit tests
	@echo "Testing with timex_disable tag (only root package)"
	@go test -tags=timex_disable -race .
	@echo "Testing without timex_disable tag (normal)"
	@go test -coverprofile=coverage.out -covermode=atomic -race ./...

lint:
	@golangci-lint run

fmt: ## Format files
	@goimports -w $$(find . -name "*.go" -not -path "./vendor/*")

benchmark: ## Run benchmarks
	@echo "Benchmarks with timex_disable tag"
	@go test -run=NONE -benchmem -benchtime=5s -bench=. -tags=timex_disable .
	@echo "Benchmarks without timex_disable tag (normal)"
	@go test -run=NONE -benchmem -benchtime=5s -bench=. .	
