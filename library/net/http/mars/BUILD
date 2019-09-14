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
        "//library/net/http/mars/binding:all-srcs",
        "//library/net/http/mars/middleware/antispam:all-srcs",
        "//library/net/http/mars/middleware/auth:all-srcs",
        "//library/net/http/mars/render:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
