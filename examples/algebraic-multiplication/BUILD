load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["main.go"],
    importpath = "github.com/marpaia/quanto/examples/algebraic-multiplication",
    visibility = ["//visibility:private"],
    deps = ["//algebra:go_default_library"],
)

go_binary(
    name = "algebraic-multiplication",
    embed = [":go_default_library"],
    visibility = ["//visibility:public"],
)
