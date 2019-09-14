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
        "feedback.go",
        "service.go",
    ],
    importpath = "valerian/app/interface/feedback/service",
    deps = [
        "//app/interface/feedback/conf:go_default_library",
        "//app/interface/feedback/dao:go_default_library",
        "//app/interface/feedback/model:go_default_library",
        "//library/database/sqalx:go_default_library",
        "//library/ecode:go_default_library",
        "//library/gid:go_default_library",
        "//library/log:go_default_library",
        "//library/net/metadata:go_default_library",
        "//vendor/github.com/jinzhu/copier:go_default_library",
    ],
)
