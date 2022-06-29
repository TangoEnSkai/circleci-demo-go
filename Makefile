.PHONY: dependencies
dependencies:
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.46.2
	@go mod download
	@go mod tidy

.PHONY: lint
lint:
	@golangci-lint run ./...

.PHONY: lint-fix
lint-fix:
	@golangci-lint run --fix ./...

.PHONY: test
test:
	@go test -race -cover -count 1 ./...

.PHONY: coverage
coverage:
	@go test -race -count 1 -coverpkg=./... -coverprofile=coverage.txt ./...