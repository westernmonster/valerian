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
        "service.go",
        "valcode.go",
    ],
    importpath = "valerian/app/interface/file/service",
    deps = [
        "//app/interface/file/conf:go_default_library",
        "//app/interface/file/model:go_default_library",
        "//library/gid:go_default_library",
        "//library/log:go_default_library",
    ],
)
