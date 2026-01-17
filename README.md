# TDD with Go

DO [Learn Go with tests](https://quii.gitbook.io/learn-go-with-tests)

## Environments

- mise 2025.10.4
- go 1.25.6
- golangci-lint v2.8.0

## Setup

```sh
# Install golangci-lint
curl -sSfL https://golangci-lint.run/install.sh | sh -s -- -b $(go env GOPATH)/bin v2.8.0
```

## Run test

### GO FUNDAMENTALS

- [Hello, world!](https://andmorefine.gitbook.io/learn-go-with-tests/go-fundamentals/hello-world)

```sh
go test ./greeting
```

- [Integers](https://andmorefine.gitbook.io/learn-go-with-tests/go-fundamentals/integers)

```sh
go test ./integers
go test -v ./integers

# See example with below command:
godoc -http :8000
# godoc: http://localhost:8000/pkg/github.com/nukopy/tdd-with-go/integers/
```

- Iteration

```sh
go test ./iteration
```

## References

- [Learn Go with tests](https://andmorefine.gitbook.io/learn-go-with-tests)
- [Go by Example](https://gobyexample.com)
- [github.com/traPtitech/traQ](https://github.com/traPtitech/traQ)
  - traQ - traP Internal Messenger Application Backend
- [github.com/traPtitech/go-traq](https://github.com/traPtitech/go-traq)
  - Go client library for the traQ API
