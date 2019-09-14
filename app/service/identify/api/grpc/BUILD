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
    name = "api_proto",
    srcs = ["pb.proto"],
    deps = ["//github.com/gogo/protobuf/gogoproto:gogoproto_proto"],
)

go_proto_library(
    name = "api_go_proto",
    compilers = ["@io_bazel_rules_go//proto:go_grpc"],
    importpath = "valerian/app/service/identify/api/grpc",
    proto = ":api_proto",
    deps = ["//github.com/gogo/protobuf/gogoproto:go_default_library"],
)

go_library(
    name = "go_default_library",
    srcs = [
        "client.go",
        "string.go",
    ],
    embed = [":api_go_proto"],
    importpath = "valerian/app/service/identify/api/grpc",
    deps = [
        "//library/net/rpc/warden:go_default_library",
        "//vendor/github.com/gogo/protobuf/proto:go_default_library",
        "@org_golang_google_grpc//:go_default_library",
    ],
)
