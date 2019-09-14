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
    srcs = [
        "dao.go",
        "mc.go",
    ],
    importpath = "valerian/app/interface/valcode/dao",
    deps = [
        "//app/interface/valcode/conf:go_default_library",
        "//library/cache/memcache:go_default_library",
        "//library/log:go_default_library",
    ],
)
