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
    name = "types_proto",
    srcs = ["status.proto"],
    deps = ["@com_google_protobuf//:any_proto"],
)

go_proto_library(
    name = "types_go_proto",
    importpath = "valerian/library/ecode/internal/types",
    proto = ":types_proto",
)

go_library(
    name = "go_default_library",
    embed = [":types_go_proto"],
    importpath = "valerian/library/ecode/internal/types",
)
