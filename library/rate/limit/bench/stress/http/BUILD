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
    srcs = ["http.go"],
    importpath = "valerian/library/rate/limit/bench/stress/http",
    deps = [
        "//library/log:go_default_library",
        "//library/net/http/mars:go_default_library",
        "//library/rate:go_default_library",
        "//library/rate/limit:go_default_library",
        "//library/rate/limit/bench/stress/conf:go_default_library",
        "//library/rate/limit/bench/stress/service:go_default_library",
        "//library/rate/vegas:go_default_library",
    ],
)
