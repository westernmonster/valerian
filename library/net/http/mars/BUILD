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
        "//library/net/http/mars/binding:all-srcs",
        "//library/net/http/mars/middleware/antispam:all-srcs",
        "//library/net/http/mars/middleware/auth:all-srcs",
        "//library/net/http/mars/render:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)

go_library(
    name = "go_default_library",
    srcs = [
        "client.go",
        "client_trace.go",
        "context.go",
        "cors.go",
        "csrf.go",
        "device.go",
        "logger.go",
        "metadata.go",
        "perf.go",
        "prometheus.go",
        "recovery.go",
        "routergroup.go",
        "server.go",
        "status_code_tracker.go",
        "trace.go",
        "utils.go",
    ],
    importpath = "valerian/library/net/http/mars",
    deps = [
        "//library/conf/dsn:go_default_library",
        "//library/conf/env:go_default_library",
        "//library/ecode:go_default_library",
        "//library/log:go_default_library",
        "//library/net/http/mars/binding:go_default_library",
        "//library/net/http/mars/render:go_default_library",
        "//library/net/ip:go_default_library",
        "//library/net/metadata:go_default_library",
        "//library/net/netutil/breaker:go_default_library",
        "//library/stat:go_default_library",
        "//library/time:go_default_library",
        "//library/tracing:go_default_library",
        "//vendor/github.com/gogo/protobuf/proto:go_default_library",
        "//vendor/github.com/gogo/protobuf/types:go_default_library",
        "//vendor/github.com/opentracing/opentracing-go:go_default_library",
        "//vendor/github.com/opentracing/opentracing-go/ext:go_default_library",
        "//vendor/github.com/opentracing/opentracing-go/log:go_default_library",
        "//vendor/github.com/pkg/errors:go_default_library",
        "//vendor/github.com/prometheus/client_golang/prometheus/promhttp:go_default_library",
        "//vendor/go.uber.org/zap:go_default_library",
        "//vendor/go.uber.org/zap/zapcore:go_default_library",
    ],
)
