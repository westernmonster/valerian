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
        "//library/email/tmpl:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)

go_library(
    name = "go_default_library",
    srcs = ["email.go"],
    importpath = "valerian/library/email",
    deps = [
        "//library/email/tmpl:go_default_library",
        "//library/email/tmpl/layouts:go_default_library",
        "//library/tracing:go_default_library",
        "//vendor/github.com/aliyun/alibaba-cloud-sdk-go/sdk:go_default_library",
        "//vendor/github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests:go_default_library",
        "//vendor/github.com/opentracing/opentracing-go:go_default_library",
        "//vendor/github.com/opentracing/opentracing-go/ext:go_default_library",
        "//vendor/github.com/sirupsen/logrus:go_default_library",
        "//vendor/github.com/ztrue/tracerr:go_default_library",
    ],
)
