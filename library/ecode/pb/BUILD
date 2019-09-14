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
    srcs = [":package-srcs"],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)

proto_library(
    name = "ecode_proto",
    srcs = ["ecode.proto"],
)

go_proto_library(
    name = "ecode_go_proto",
    importpath = "valerian/library/ecode/pb",
    proto = ":ecode_proto",
)

go_library(
    name = "go_default_library",
    srcs = ["ecode.go"],
    embed = [":ecode_go_proto"],
    importpath = "valerian/library/ecode/pb",
    deps = [
        "//library/ecode:go_default_library",
        "@io_bazel_rules_go//proto/wkt:any_go_proto",
    ],
)
