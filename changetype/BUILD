load("@io_bazel_rules_go//go:def.bzl", "go_test")

go_test(
    name = "go_default_test",
    size = "small",
    srcs = ["compat_test.go"],
    deps = [
        "//:go_default_library",
        "//changetype/a:go_default_library",
        "//changetype/b:go_default_library",
    ],
)
