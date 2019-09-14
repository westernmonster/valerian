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
        "http.go",
        "init.go",
        "mutex.go",
        "mux.go",
        "noop.go",
        "tag.go",
        "trace.go",
    ],
    importpath = "valerian/library/tracing",
    deps = [
        "//library/conf/dsn:go_default_library",
        "//library/conf/env:go_default_library",
        "//library/log:go_default_library",
        "//library/time:go_default_library",
        "//vendor/github.com/opentracing-contrib/go-stdlib/nethttp:go_default_library",
        "//vendor/github.com/opentracing/opentracing-go:go_default_library",
        "//vendor/github.com/opentracing/opentracing-go/log:go_default_library",
        "//vendor/github.com/pkg/errors:go_default_library",
        "//vendor/github.com/uber/jaeger-client-go/config:go_default_library",
        "//vendor/github.com/uber/jaeger-client-go/rpcmetrics:go_default_library",
        "//vendor/github.com/uber/jaeger-lib/metrics:go_default_library",
        "//vendor/github.com/uber/jaeger-lib/metrics/prometheus:go_default_library",
        "//vendor/go.uber.org/zap:go_default_library",
    ],
)
