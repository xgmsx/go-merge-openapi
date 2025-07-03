export

.PHONY: lint
lint:
	golangci-lint run --fix ./...

.PHONY: fmt
fmt:
	gofumpt -extra -w .
	gci write -s standard -s default -s "prefix(github.com/xgmsx/go-merge-openapi)" .

.PHONY: build
build:
	go build .

.PHONY: run
run:
	go run . -config ./examples/api/petstore.openapi.yaml -output ./petstore.openapi.yaml

.PHONY: cov
cov:
	go test -cover -coverprofile=coverage.txt ./...
	go tool cover -html=coverage.txt -o coverage.html
	go tool cover -func=coverage.txt

.PHONY: test
test:
	go test ./...
