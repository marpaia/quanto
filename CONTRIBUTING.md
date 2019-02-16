# Contributor Guide

The contributor guide has the following sections:

- [Setup](#setup)
- [Build](#build)
- [Test](#test)
- [Develop](#develop)

## Setup

To get your development environment setup, you first must clone the source repository:

```
git clone git@github.com:marpaia/math-go.git
```

To build and test the code, it is required to have the following minimal toolset installed:

- [The Go programming language](https://golang.org/dl/)

The following tools are not required, but recommended:

- [The Bazel build system](https://docs.bazel.build/versions/master/install.html)

## Build

While this repository can be built with the `go` command (`go build`, `go test`, `go mod`, etc), there is also support for building with [Bazel](https://www.bazel.build/). When the `go` command does what you need, feel free to use it. To run certain build tasks that are outside of the `go` tool's capabilities, of course one must use `bazel`.

From the root of the repository, you can run the following to build everything using `bazel`:

```
bazel build ...
```

For this, you can also use the `go` command directly to just build the source code:

```
go build ./...
```

To build a specific target in a more minimal way, you can specify an individual package:

```
bazel build //calculus/...
```

Similarly, the analagous `go` command is:

```
go build ./calculus/...
```

## Test

From the root of the repository, you can run the following to run all of the tests:

```
bazel test ...
```

With the `go` command, you should run:

```
go test ./...
```

To run a specific test in a more minimal way, you can specify an individual package:

```
bazel test //calculus/...
```

Or with the `go` command:

```
go test ./calculus/...
```

### Running A Full CI Build Locally

The CircleCI configuration includes a number of lint and test steps. If you'd like to run a complete, representative CI build locally, download the `circleci` CLI tool. See the [official installation instructions](https://circleci.com/docs/2.0/local-cli/#installing-the-circleci-local-cli-on-macos-and-linux-distros) for download information.

Once you have the tool installed in your path, run the following from the root of the repository:

```
circleci build
```

## Develop

### Bazel Configurations

#### Generating Configs

Because we follow the patterns required by the `go` tool, we can use the [Gazelle](https://github.com/bazelbuild/bazel-gazelle) project to automatically generate Bazel configurations. This allows us to be in a best-of-both-worlds situation where we are able to auto-generate anything that seems redundant and take advantage of Bazel for reproducible builds, extensions, task running, building container images, etc.

If you add new files, dependencies, libraries, or commands, you will likely need to regenerate the Bazel `BUILD` files. You can generate the Bazel `BUILD` files by running the following from the root of the repository:

```
bazel run //:gazelle
```

Bazel manages all dependencies internally with individual `go_repository` stanzas in the project `WORKSPACE`. Since the `go` tool will maintain a manifest similar to this in `go.mod` and `go.sum`, we can use these files to automatically maintain the `WORKSPACE` file. To generate the `go_repository` definitions in the project `WORKSPACE`, you can run the following:

```
bazel run //:gazelle -- update-repos -from_file=go.mod
```
#### Formatting Configs

To format all Bazel `BUILD` files using [`buildifier`](https://github.com/bazelbuild/buildtools/tree/master/buildifier), run the following:

```
bazel run //:buildifier
```
