[![Build Status](https://travis-ci.com/savannahghi/feedlib.svg?branch=main)](https://travis-ci.com/savannahghi/feedlib)
[![Maintained](https://img.shields.io/badge/Maintained-Actively-informational.svg?style=for-the-badge)](https://shields.io/)
[![MIT license](https://img.shields.io/badge/license-MIT-brightgreen.svg)](https://opensource.org/licenses/MIT) ![Linting and Tests](https://github.com/savannahghi/feedlib/actions/workflows/ci.yml/badge.svg)
[![Coverage Status](https://coveralls.io/repos/github/savannahghi/feedlib/badge.svg?branch=main)](https://coveralls.io/github/savannahghi/feedlib?branch=main)
# Feed Library
feedlib  is an open source project — it's one among many other shared libraries that make up the wider ecosystem of software made and open sourced by Savannah Informatics Limited.

A shared library for Be.Well Golang services that is responsible for rendering user-feed and engagement.

### Installing it
feedlib is compatible with modern Go releases in module mode, with Go installed:

```bash
go get -u github.com/savannahghi/feedlib

```
will resolve and add the package to the current development module, along with its dependencies.

Alternatively the same can be achieved if you use import in a package:

```go
import "github.com/savannahghi/feedlib"

```
and run `go get` without parameters.

The package name is `feedlib`


### Developing

The default branch library is `main`

We try to follow semantic versioning ( <https://semver.org/> ). For that reason,
every major, minor and point release should be _tagged_.

```
git tag -m "v0.0.1" "v0.0.1"
git push --tags
```

Continuous integration tests *must* pass on Travis CI. Our coverage threshold
is 90% i.e you *must* keep coverage above 90%.


## Environment variables

In order to run tests, you need to have an `env.sh` file similar to this one:

```bash
# Application settings
export SCHEMA_HOST=<optional>

```

This file *must not* be committed to version control.

It is important to _export_ the environment variables. If they are not exported,
they will not be visible to child processes e.g `go test ./...`.

These environment variables should also be set up on Travis CI environment variable section.

## Contributing ##
Contributions are of course always welcome. The calling pattern is pretty well established, so adding new methods is relatively
straightforward. See [`CONTRIBUTING.md`](CONTRIBUTING.md) for details.

## Versioning ##

In general, feedlib follows [semver](https://semver.org/) as closely as we
can for tagging releases of the package. For self-contained libraries, the
application of semantic versioning is relatively straightforward and generally
understood. We've adopted the following
versioning policy:

* We increment the **major version** with any incompatible change to
	non-preview functionality, including changes to the exported Go API surface
	or behavior of the API.
* We increment the **minor version** with any backwards-compatible changes to
	functionality, as well as any changes to preview functionality in the GitHub
	API. GitHub makes no guarantee about the stability of preview functionality,
	so neither do we consider it a stable part of the go-github API.
* We increment the **patch version** with any backwards-compatible bug fixes.

## License ##

This library is distributed under the MIT license found in the [LICENSE](./LICENSE)
file.