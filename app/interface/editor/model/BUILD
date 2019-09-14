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
        "arg.go",
        "resp.go",
    ],
    importpath = "valerian/app/interface/editor/model",
    deps = [
        "//vendor/github.com/go-ozzo/ozzo-validation:go_default_library",
        "//vendor/github.com/go-ozzo/ozzo-validation/is:go_default_library",
    ],
)
