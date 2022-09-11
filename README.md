# go-set

[![actions-workflow-test][actions-workflow-test-badge]][actions-workflow-test]
[![release][release-badge]][release]
[![pkg.go.dev][pkg.go.dev-badge]][pkg.go.dev]
[![license][license-badge]][license]

An experimental Go package that defines various methods for a [set](https://en.wikipedia.org/wiki/Set_(abstract_data_type)) implemented with Go generics.

## Usage

See the [examples](./example_test.go).

## Alternatives

- [golang.org/x/exp/slices](https://pkg.go.dev/golang.org/x/exp/slices)
  - `go-set` may be a better option if you don't find functions you want in this packages.
- [k8s.io/apimachinery/pkg/util/sets](https://pkg.go.dev/k8s.io/apimachinery/pkg/util/sets)
  - `go-set` may be a better option if you prefer to use generics.

<!-- badge links -->

[actions-workflow-test]: https://github.com/micnncim/go-set/actions?query=workflow%3ATest
[actions-workflow-test-badge]: https://img.shields.io/github/workflow/status/micnncim/go-set/Test?label=Test&style=for-the-badge&logo=github

[release]: https://github.com/micnncim/go-set/releases
[release-badge]: https://img.shields.io/github/v/release/micnncim/go-set?style=for-the-badge&logo=github

[pkg.go.dev]: https://pkg.go.dev/github.com/micnncim/go-set?tab=overview
[pkg.go.dev-badge]: http://bit.ly/pkg-go-dev-badge

[license]: LICENSE
[license-badge]: https://img.shields.io/github/license/micnncim/go-set?style=for-the-badge
