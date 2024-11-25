## Playground

https://go.dev/play

## Import

### Import from local folder

- run in project root folder
  - `go mod init projectname`
- use import
  - `import "projectname/folder"`

## Slice

https://ueokande.github.io/go-slice-tricks

## Module

- init module
  `go mod init golearning`
- import module package
  `import golearning/7-scope/lib`

### Local modules

- create main module

```go
module main

go 1.23.3

replace golearning => ../golearning
```

- run `go get golearning`

### External modules

`package.json` === `go.mod`  
`package-lock.json` === `go.sum`

```go
module somemodule

go 1.23

require github.com/stretchr/testify v1.7.0
```

#### Vendor

local copy of all dependencies

`go mod vendor` - create vendors

## Documentation

- `go install golang.org/x/tools/...@latest`
- `godoc -http=:6060`
- open `localhost:6060`

## Unit tests

There's a catch. This works well if:

- `foo.go` is in package `foo`
- `foo_test.go` is in package `foo_test` and imports `foo`.

If `foo_test.go` and `foo.go` are the same package (a common case) then you must name all other files required to build `foo_test`

`$ go test foo_test.go foo.go`

### Coverage

```bash
go test file.go -coverprofile=cover.out
go tool cover -html=cover.out
```

### Windows firewall

- go to Windows Defender Firewall
- select `Inbound Rules`
- click `New Rule`
  - choose `Port`
  - select `TCP`
  - `8080`
  - `Allow the connection`
