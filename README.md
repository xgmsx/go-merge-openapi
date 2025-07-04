# go-merge-openapi

[![Linting](https://github.com/xgmsx/go-merge-openapi/actions/workflows/golangci-lint.yml/badge.svg?branch=main)](https://github.com/xgmsx/go-merge-openapi/actions/workflows/golangci-lint.yml)
[![Tests](https://github.com/xgmsx/go-merge-openapi/actions/workflows/coverage.yml/badge.svg?branch=main)](https://github.com/xgmsx/go-merge-openapi/actions/workflows/coverage.yml)
[![CodeQL](https://github.com/xgmsx/go-merge-openapi/actions/workflows/codeql.yml/badge.svg?branch=main)](https://github.com/xgmsx/go-merge-openapi/actions/workflows/codeql.yml)
[![Coverage_Report](https://img.shields.io/badge/Coverage_Report-50.5%25-yellow)](https://xgmsx.github.io/go-merge-openapi)

### Table of Contents
- [Installation](#installation)
- [Usage](#usage)
- [Examples](#examples)

### Installation

```shell
go install github.com/xgmsx/go-merge-openapi@1.0.0

go install github.com/xgmsx/go-merge-openapi@latest
```

### Usage

```shell
go-merge-openapi -config ./examples/api/petstore.openapi.yaml -output ./petstore.openapi.yaml

go-merge-openapi -config ./examples/api/petstore.openapi.yaml -output ./petstore.openapi.json
```

### Examples

Input `examples/api/petstore.openapi.yaml` file:
```yaml
openapi: 3.0.4
info:
  title: Swagger Petstore - OpenAPI 3.0
  version: 1.0.12
paths:
  /api/v1/pet/{petId}:
    $ref: ./paths/pet_by_id.yaml
```

Output `petstore.openapi.merged.yaml` file:
```yaml
openapi: 3.0.4
info:
  title: Swagger Petstore - OpenAPI 3.0
  version: 1.0.12
paths:
  /api/v1/pet/{petId}:
    get:
      tags:
        - pet
      summary: Find pet by ID.
      description: Returns a single pet.
      operationId: getPetById
      parameters:
        - name: petId
          in: path
          description: ID of pet that needs to be updated
          required: true
          schema:
            type: integer
            format: int64
      responses:
        '200':
          description: Successful operation
          content:
            application/json:
              schema:
                type: object
                properties:
                  id:
                    type: integer
                    format: int64
                  name:
                    type: string
                  status:
                    type: string
        '400':
          description: Invalid ID supplied
        '404':
          description: Pet not found
```