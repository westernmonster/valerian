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
        "country_cde.go",
        "service.go",
    ],
    importpath = "valerian/app/interface/location/service",
    deps = [
        "//app/interface/location/conf:go_default_library",
        "//app/interface/location/dao:go_default_library",
        "//app/interface/location/model:go_default_library",
        "//library/database/sqalx:go_default_library",
        "//library/log:go_default_library",
    ],
)
