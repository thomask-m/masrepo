load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "rle",
    srcs = ["rle.go"],
    importpath = "github.com/example/project/go/src/mas/compression/rle",
    visibility = ["//visibility:public"],
)

go_test(
    name = "rle_test",
    srcs = ["rle_test.go"],
    embed = [":rle"],
    deps = ["@com_github_stretchr_testify//assert:rle"],
)
