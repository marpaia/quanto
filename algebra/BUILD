load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "algebra.go",
        "arithmetic.go",
    ],
    importpath = "github.com/marpaia/quanto/algebra",
    visibility = ["//visibility:public"],
)

go_test(
    name = "go_default_test",
    srcs = ["algebra_test.go"],
    embed = [":go_default_library"],
    deps = ["//internal/assert:go_default_library"],
)
