.PHONY: dependencies
dependencies:
	@go mod vendor
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
	@go test -v -race -cover -count 1 ./...

.PHONY: coverage
coverage:
	@go test -race -count 1 -coverpkg=./... -coverprofile=coverage.txt ./...

.PHONY: deploy
deploy:
	@echo deploying to gcp...
	@echo deployed!
