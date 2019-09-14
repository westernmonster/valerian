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
    srcs = [
        ":package-srcs",
        "//vendor/github.com/opentracing/opentracing-go/ext:all-srcs",
        "//vendor/github.com/opentracing/opentracing-go/log:all-srcs",
        "//vendor/github.com/opentracing/opentracing-go/mocktracer:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
