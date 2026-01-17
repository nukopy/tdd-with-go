# TDD with Go

[![CI][github-actions-badge]][github-actions-url]
[![codecov][codecov-badge]][codecov-url]

[github-actions-badge]: https://github.com/nukopy/tdd-with-go/actions/workflows/ci.yml/badge.svg
[github-actions-url]: https://github.com/nukopy/tdd-with-go/actions/workflows/ci.yml
[codecov-badge]: https://codecov.io/github/nukopy/tdd-with-go/graph/badge.svg?token=x92k8kBb3t
[codecov-url]: https://codecov.io/github/nukopy/tdd-with-go

DO [Learn Go with tests](https://quii.gitbook.io/learn-go-with-tests)

## Environments

- mise 2025.10.4
- go 1.25.6
- golangci-lint v2.8.0

## Setup

- Installation of golangci-lint

```sh
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
