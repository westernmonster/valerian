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
