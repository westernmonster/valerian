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
    name = "render_proto",
    srcs = ["pb.proto"],
    deps = [
        "//github.com/gogo/protobuf/gogoproto:gogoproto_proto",
        "@com_google_protobuf//:any_proto",
    ],
)

go_proto_library(
    name = "render_go_proto",
    importpath = "valerian/library/net/http/mars/render",
    proto = ":render_proto",
    deps = ["//github.com/gogo/protobuf/gogoproto:go_default_library"],
)

go_library(
    name = "go_default_library",
    srcs = [
        "data.go",
        "json.go",
        "protobuf.go",
        "redirect.go",
        "render.go",
        "string.go",
        "xml.go",
    ],
    embed = [":render_go_proto"],
    importpath = "valerian/library/net/http/mars/render",
    deps = [
        "//vendor/github.com/gogo/protobuf/proto:go_default_library",
        "//vendor/github.com/pkg/errors:go_default_library",
    ],
)
