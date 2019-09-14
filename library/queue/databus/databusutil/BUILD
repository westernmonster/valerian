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
        "group.go",
    ],
    importpath = "valerian/library/queue/databus/databusutil",
    deps = [
        "//library/queue/databus:go_default_library",
        "//library/time:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = ["group_test.go"],
    embed = [":go_default_library"],
    deps = [
        "//library/log:go_default_library",
        "//library/queue/databus:go_default_library",
        "//library/time:go_default_library",
        "//vendor/golang.org/x/sync/errgroup:go_default_library",
    ],
)
