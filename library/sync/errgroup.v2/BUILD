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
    srcs = [":package-srcs"],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)

go_library(
    name = "go_default_library",
    srcs = [
        "doc.go",
        "errgroup.go",
    ],
    importpath = "valerian/library/sync/errgroup.v2",
)

go_test(
    name = "go_default_test",
    srcs = [
        "errgroup_test.go",
        "example_test.go",
    ],
    embed = [":go_default_library"],
)
