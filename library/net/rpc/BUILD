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
        "//library/net/rpc/context:all-srcs",
        "//library/net/rpc/interceptor:all-srcs",
        "//library/net/rpc/warden:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
