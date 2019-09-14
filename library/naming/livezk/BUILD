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

go_library(
    name = "go_default_library",
    srcs = ["livezk.go"],
    importpath = "valerian/library/naming/livezk",
    deps = [
        "//library/log:go_default_library",
        "//library/naming:go_default_library",
        "//library/time:go_default_library",
        "//vendor/github.com/samuel/go-zookeeper/zk:go_default_library",
    ],
)
