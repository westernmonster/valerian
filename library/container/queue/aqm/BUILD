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
    srcs = ["codel.go"],
    importpath = "valerian/library/container/queue/aqm",
    deps = ["//library/ecode:go_default_library"],
)

go_test(
    name = "go_default_test",
    srcs = ["codel_test.go"],
    embed = [":go_default_library"],
    deps = ["//library/ecode:go_default_library"],
)
