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
        "locale.go",
        "locale-packr.go",
    ],
    importpath = "valerian/library/locale",
    deps = [
        "//vendor/github.com/gobuffalo/packr:go_default_library",
        "//vendor/github.com/nicksnyder/go-i18n/v2/i18n:go_default_library",
        "//vendor/golang.org/x/text/language:go_default_library",
        "//vendor/gopkg.in/yaml.v2:go_default_library",
    ],
)
