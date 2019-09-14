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
    srcs = [
        "arg_rpc.go",
        "client.go",
        "const.go",
        "model.go",
    ],
    importpath = "valerian/app/infra/config/model",
    deps = [
        "//library/database/sqlx/types:go_default_library",
        "//library/time:go_default_library",
    ],
)
