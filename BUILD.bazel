load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")
load("@bazel_gazelle//:def.bzl", "gazelle")

gazelle(
    name = "gazelle",
    prefix = "github.com/prysmaticlabs/beacon-chain",
)

go_library(
    name = "go_default_library",
    srcs = ["main.go"],
    importpath = "github.com/prysmaticlabs/beacon-chain",
    visibility = ["//visibility:private"],
)

go_binary(
    name = "beacon-chain",
    embed = [":go_default_library"],
    visibility = ["//visibility:public"],
)
