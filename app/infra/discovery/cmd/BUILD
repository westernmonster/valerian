package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_binary",
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
    srcs = ["main.go"],
    importpath = "valerian/app/infra/discovery/cmd",
    deps = [
        "//app/infra/discovery/conf:go_default_library",
        "//app/infra/discovery/http:go_default_library",
        "//app/infra/discovery/service:go_default_library",
        "//library/conf/env:go_default_library",
        "//library/log:go_default_library",
        "//library/tracing:go_default_library",
        "//vendor/github.com/spf13/pflag:go_default_library",
    ],
)

go_binary(
    name = "cmd",
    embed = [":go_default_library"],
)
