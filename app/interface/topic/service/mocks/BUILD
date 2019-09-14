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
    srcs = ["IDao.go"],
    importpath = "valerian/app/interface/topic/service/mocks",
    deps = [
        "//app/interface/topic/model:go_default_library",
        "//library/database/sqalx:go_default_library",
        "//vendor/github.com/stretchr/testify/mock:go_default_library",
    ],
)
