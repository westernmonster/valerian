package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
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
    srcs = ["docs.go"],
    importpath = "valerian/app/interface/topic/model/docs",
    deps = [
        "//vendor/github.com/alecthomas/template:go_default_library",
        "//vendor/github.com/swaggo/swag:go_default_library",
    ],
)
