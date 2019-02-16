load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["main.go"],
    importpath = "github.com/marpaia/quanto/examples/vector-arithmetic",
    visibility = ["//visibility:private"],
    deps = ["//vector:go_default_library"],
)

go_binary(
    name = "vector-arithmetic",
    embed = [":go_default_library"],
    visibility = ["//visibility:public"],
)
