load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_prefix")

go_prefix("github.com/kyleconroy/protocompat")

go_library(
    name = "go_default_library",
    srcs = ["protocompat.go"],
    visibility = ["//visibility:public"],
    deps = ["//vendor/github.com/golang/protobuf/proto:go_default_library"],
)
