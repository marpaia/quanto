load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = ["geometry.go"],
    importpath = "github.com/marpaia/quanto/geometry",
    visibility = ["//visibility:public"],
    deps = ["//vector:go_default_library"],
)

go_test(
    name = "go_default_test",
    srcs = ["geometry_test.go"],
    embed = [":go_default_library"],
    deps = ["//vector:go_default_library"],
)
