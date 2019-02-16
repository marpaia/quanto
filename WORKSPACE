workspace(name = "com_github_quanto_quanto")

################################################################################
# General Initialization
################################################################################

# The native http_archive rule is deprecated. This is a drop-in replacement.
load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

http_archive(
    name = "com_github_bazelbuild_buildtools",
    strip_prefix = "buildtools-0.22.0",
    url = "https://github.com/bazelbuild/buildtools/archive/0.22.0.zip",
)

load("@com_github_bazelbuild_buildtools//buildifier:deps.bzl", "buildifier_dependencies")

buildifier_dependencies()

################################################################################
# Go Build Definitions
################################################################################

# Download, load, and initialize the Go build rules.
http_archive(
    name = "io_bazel_rules_go",
    strip_prefix = "rules_go-0.17.0",
    urls = ["https://github.com/bazelbuild/rules_go/archive/0.17.0.zip"],
)

load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")

go_rules_dependencies()

go_register_toolchains()

# Download, load, and initialize the Gazelle tool for generating BUILD files for
# Go code.
http_archive(
    name = "bazel_gazelle",
    strip_prefix = "bazel-gazelle-0.16.0",
    urls = ["https://github.com/bazelbuild/bazel-gazelle/archive/0.16.0.zip"],
)

load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies", "go_repository")

gazelle_dependencies()

################################################################################
# Go Dependencies
#
# The gazelle tool will automatically write Go dependencies in this section
# based on the `go.mod` file. For that reason, this section should always be at
# the end of this file.
################################################################################
