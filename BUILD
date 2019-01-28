load("@bazel_gazelle//:def.bzl", "gazelle")

# gazelle:prefix github.com/marpaia/math-go
gazelle(
    name = "gazelle",
    extra_args = ["--build_file_name=BUILD"],
)

load("@com_github_bazelbuild_buildtools//buildifier:def.bzl", "buildifier")

buildifier(name = "buildifier")
