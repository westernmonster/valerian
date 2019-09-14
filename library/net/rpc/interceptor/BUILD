package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_test",
    "go_library",
)

filegroup(
    name = "package-srcs",
    srcs = glob(["**"]),
    tags = ["automanaged"],
    visibility = ["//visibility:private"],
)

filegroup(
    name = "all-srcs",
    srcs = [":package-srcs"],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)

go_library(
    name = "go_default_library",
    srcs = ["interceptor.go"],
    importpath = "valerian/library/net/rpc/interceptor",
    deps = [
        "//library/ecode:go_default_library",
        "//library/log:go_default_library",
        "//library/net/rpc/context:go_default_library",
        "//library/stat:go_default_library",
        "//vendor/go.uber.org/zap:go_default_library",
        "//vendor/golang.org/x/time/rate:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = ["interceptor_test.go"],
    embed = [":go_default_library"],
    deps = [
        "//library/ecode:go_default_library",
        "//library/net/rpc/context:go_default_library",
        "//vendor/golang.org/x/time/rate:go_default_library",
    ],
)
