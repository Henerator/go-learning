## Playground

https://go.dev/play

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
