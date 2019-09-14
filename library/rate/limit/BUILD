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
    srcs = [
        ":package-srcs",
        "//library/rate/limit/bench/stress/cmd:all-srcs",
        "//library/rate/limit/bench/stress/conf:all-srcs",
        "//library/rate/limit/bench/stress/dao:all-srcs",
        "//library/rate/limit/bench/stress/http:all-srcs",
        "//library/rate/limit/bench/stress/service:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
