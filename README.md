# go-merge-openapi

[![Linting](https://github.com/xgmsx/go-merge-openapi/actions/workflows/golangci-lint.yml/badge.svg?branch=main)](https://github.com/xgmsx/go-merge-openapi/actions/workflows/golangci-lint.yml)
[![Tests](https://github.com/xgmsx/go-merge-openapi/actions/workflows/coverage.yml/badge.svg?branch=main)](https://github.com/xgmsx/go-merge-openapi/actions/workflows/coverage.yml)
[![CodeQL](https://github.com/xgmsx/go-merge-openapi/actions/workflows/codeql.yml/badge.svg?branch=main)](https://github.com/xgmsx/go-merge-openapi/actions/workflows/codeql.yml)
[![Coverage_Report](https://img.shields.io/badge/Coverage_Report-50.5%25-yellow)](https://xgmsx.github.io/go-merge-openapi)

### Installation

```shell
go install github.com/xgmsx/go-merge-openapi@latest
```

### Usage

```shell
go-merge-openapi -config ./examples/api/petstore.openapi.yaml -output ./petstore.openapi.yaml

go-merge-openapi -config ./examples/api/petstore.openapi.yaml -output ./petstore.openapi.json
```
