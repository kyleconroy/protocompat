load("@io_bazel_rules_go//go:def.bzl", "go_test")

go_test(
    name = "go_default_test",
    size = "small",
    srcs = ["compat_test.go"],
    deps = [
        "//:go_default_library",
        "//addfield/a:go_default_library",
        "//addfield/b:go_default_library",
    ],
)
