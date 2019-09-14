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
        "//library/ecode/internal/types:all-srcs",
        "//library/ecode/pb:all-srcs",
        "//library/ecode/tip:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)

go_library(
    name = "go_default_library",
    srcs = [
        "common_ecode.go",
        "ecode.go",
        "status.go",
    ],
    importpath = "valerian/library/ecode",
    deps = [
        "//library/ecode/internal/types:go_default_library",
        "//vendor/github.com/pkg/errors:go_default_library",
        "@com_github_golang_protobuf//proto:go_default_library",
        "@com_github_golang_protobuf//ptypes:go_default_library_gen",
    ],
)

go_test(
    name = "go_default_test",
    srcs = [
        "ecode_test.go",
        "status_test.go",
    ],
    embed = [":go_default_library"],
    deps = [
        "//library/ecode/internal/types:go_default_library",
        "//vendor/github.com/smartystreets/goconvey/convey:go_default_library",
        "//vendor/github.com/stretchr/testify/assert:go_default_library",
        "@io_bazel_rules_go//proto/wkt:timestamp_go_proto",
    ],
)
