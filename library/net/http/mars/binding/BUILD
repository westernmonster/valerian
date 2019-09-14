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
        "binding.go",
        "default_validator.go",
        "form.go",
        "form_mapping.go",
        "json.go",
        "query.go",
        "tags.go",
        "xml.go",
    ],
    importpath = "valerian/library/net/http/mars/binding",
    deps = [
        "//vendor/github.com/pkg/errors:go_default_library",
        "//vendor/gopkg.in/go-playground/validator.v9:go_default_library",
    ],
)
