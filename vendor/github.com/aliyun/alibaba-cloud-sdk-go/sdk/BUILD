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
        "//vendor/github.com/aliyun/alibaba-cloud-sdk-go/sdk/auth:all-srcs",
        "//vendor/github.com/aliyun/alibaba-cloud-sdk-go/sdk/endpoints:all-srcs",
        "//vendor/github.com/aliyun/alibaba-cloud-sdk-go/sdk/errors:all-srcs",
        "//vendor/github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests:all-srcs",
        "//vendor/github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses:all-srcs",
        "//vendor/github.com/aliyun/alibaba-cloud-sdk-go/sdk/utils:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
