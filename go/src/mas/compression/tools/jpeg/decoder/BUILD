load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["main.go"],
    importpath = "github.com/thomask-m/masrepo/go/src/mas/compression/tools/jpeg/decoder",
    visibility = ["//visibility:private"],
    deps = ["//go/src/mas/compression/jpeg:go_default_library"],
)

go_binary(
    name = "decoder",
    embed = [":go_default_library"],
    visibility = ["//visibility:public"],
)
