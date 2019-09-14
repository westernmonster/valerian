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
        "id_cert.go",
        "service.go",
    ],
    importpath = "valerian/app/interface/certification/service",
    deps = [
        "//app/interface/certification/conf:go_default_library",
        "//app/interface/certification/dao:go_default_library",
        "//app/interface/certification/model:go_default_library",
        "//library/cloudauth:go_default_library",
        "//library/database/sqalx:go_default_library",
        "//library/ecode:go_default_library",
        "//library/gid:go_default_library",
        "//library/log:go_default_library",
        "//library/net/metadata:go_default_library",
        "//vendor/github.com/aliyun/alibaba-cloud-sdk-go/sdk:go_default_library",
    ],
)
