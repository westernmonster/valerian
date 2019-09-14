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
    srcs = [
        ":package-srcs",
        "//library/email/tmpl/layouts:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)

go_library(
    name = "go_default_library",
    srcs = [
        "reg_valcode.qtpl.go",
        "reset_valcode.qtpl.go",
    ],
    importpath = "valerian/library/email/tmpl",
    deps = [
        "//library/email/tmpl/layouts:go_default_library",
        "//vendor/github.com/valyala/quicktemplate:go_default_library",
    ],
)
