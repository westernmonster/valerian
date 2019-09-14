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
        "//vendor/github.com/uber/jaeger-client-go/config:all-srcs",
        "//vendor/github.com/uber/jaeger-client-go/internal/baggage:all-srcs",
        "//vendor/github.com/uber/jaeger-client-go/internal/spanlog:all-srcs",
        "//vendor/github.com/uber/jaeger-client-go/internal/throttler:all-srcs",
        "//vendor/github.com/uber/jaeger-client-go/log:all-srcs",
        "//vendor/github.com/uber/jaeger-client-go/rpcmetrics:all-srcs",
        "//vendor/github.com/uber/jaeger-client-go/thrift:all-srcs",
        "//vendor/github.com/uber/jaeger-client-go/thrift-gen/agent:all-srcs",
        "//vendor/github.com/uber/jaeger-client-go/thrift-gen/baggage:all-srcs",
        "//vendor/github.com/uber/jaeger-client-go/thrift-gen/jaeger:all-srcs",
        "//vendor/github.com/uber/jaeger-client-go/thrift-gen/sampling:all-srcs",
        "//vendor/github.com/uber/jaeger-client-go/thrift-gen/zipkincore:all-srcs",
        "//vendor/github.com/uber/jaeger-client-go/transport:all-srcs",
        "//vendor/github.com/uber/jaeger-client-go/utils:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
