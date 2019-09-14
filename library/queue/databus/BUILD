load(
    "@io_bazel_rules_go//proto:def.bzl",
    "go_proto_library",
)

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
        "//library/queue/databus/databusutil:all-srcs",
        "//library/queue/databus/metadata:all-srcs",
        "//library/queue/databus/report:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)

proto_library(
    name = "databus_proto",
    srcs = ["databus.proto"],
    deps = ["//github.com/gogo/protobuf/gogoproto:gogoproto_proto"],
)

go_proto_library(
    name = "databus_go_proto",
    importpath = "valerian/library/queue/databus",
    proto = ":databus_proto",
    deps = ["//github.com/gogo/protobuf/gogoproto:go_default_library"],
)

go_library(
    name = "go_default_library",
    srcs = ["databus.go"],
    embed = [":databus_go_proto"],
    importpath = "valerian/library/queue/databus",
    deps = [
        "//library/cache/redis:go_default_library",
        "//library/conf/env:go_default_library",
        "//library/container/pool:go_default_library",
        "//library/log:go_default_library",
        "//library/net/netutil:go_default_library",
        "//library/stat/prom:go_default_library",
        "//library/time:go_default_library",
    ],
)
