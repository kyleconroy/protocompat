load("@io_bazel_rules_go//go:def.bzl", "go_prefix", "go_test")

go_prefix("github.com/kyleconroy/protocompat")

go_test(
    name = "go_default_test",
    size = "small",
    srcs = [
        "addfield_test.go",
        "changename_test.go",
        "changetag_test.go",
        "changetype_test.go",
        "compat_test.go",
        "removefield_test.go",
    ],
    deps = [
        "//one:go_default_library",
        "//two:go_default_library",
        "//vendor/github.com/golang/protobuf/proto:go_default_library",
    ],
)
