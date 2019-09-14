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
        "//vendor/github.com/smartystreets/assertions/internal/go-render/render:all-srcs",
        "//vendor/github.com/smartystreets/assertions/internal/oglematchers:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
