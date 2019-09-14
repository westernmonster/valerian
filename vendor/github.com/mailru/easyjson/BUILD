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
        "//vendor/github.com/mailru/easyjson/buffer:all-srcs",
        "//vendor/github.com/mailru/easyjson/jlexer:all-srcs",
        "//vendor/github.com/mailru/easyjson/jwriter:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
