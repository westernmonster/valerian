package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_test",
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
        "//library/net/rpc/warden/balancer/wrr:all-srcs",
        "//library/net/rpc/warden/benchmark/bench/client:all-srcs",
        "//library/net/rpc/warden/benchmark/bench/proto:all-srcs",
        "//library/net/rpc/warden/benchmark/bench/server:all-srcs",
        "//library/net/rpc/warden/benchmark/helloworld/client:all-srcs",
        "//library/net/rpc/warden/benchmark/helloworld/server:all-srcs",
        "//library/net/rpc/warden/encoding/json:all-srcs",
        "//library/net/rpc/warden/examples/client:all-srcs",
        "//library/net/rpc/warden/examples/grpcDebug:all-srcs",
        "//library/net/rpc/warden/examples/server:all-srcs",
        "//library/net/rpc/warden/metadata:all-srcs",
        "//library/net/rpc/warden/proto/testproto:all-srcs",
        "//library/net/rpc/warden/resolver:all-srcs",
        "//library/net/rpc/warden/status:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
