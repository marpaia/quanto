workspace(name = "com_github_marpaia_math_go")

################################################################################
# General Initialization
################################################################################

# The native http_archive rule is deprecated. This is a drop-in replacement.
load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

http_archive(
    name = "com_github_bazelbuild_buildtools",
    strip_prefix = "buildtools-0.20.0",
    url = "https://github.com/bazelbuild/buildtools/archive/0.20.0.zip",
)

load("@com_github_bazelbuild_buildtools//buildifier:deps.bzl", "buildifier_dependencies")

buildifier_dependencies()

################################################################################
# Go Build Definitions
################################################################################

# Download, load, and initialize the Go build rules.
http_archive(
    name = "io_bazel_rules_go",
    strip_prefix = "rules_go-0.16.6",
    urls = ["https://github.com/bazelbuild/rules_go/archive/0.16.6.zip"],
)

load("@io_bazel_rules_go//go:def.bzl", "go_register_toolchains", "go_rules_dependencies")

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

go_repository(
    name = "com_github_davecgh_go_spew",
    importpath = "github.com/davecgh/go-spew",
    tag = "v1.1.0",
)

go_repository(
    name = "com_github_pmezard_go_difflib",
    importpath = "github.com/pmezard/go-difflib",
    tag = "v1.0.0",
)

go_repository(
    name = "com_github_stretchr_objx",
    importpath = "github.com/stretchr/objx",
    tag = "v0.1.0",
)

go_repository(
    name = "com_github_stretchr_testify",
    importpath = "github.com/stretchr/testify",
    tag = "v1.3.0",
)
