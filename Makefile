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