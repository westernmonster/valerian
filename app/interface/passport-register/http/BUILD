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
        "http.go",
        "register.go",
    ],
    importpath = "valerian/app/interface/passport-register/http",
    deps = [
        "//app/interface/passport-register/conf:go_default_library",
        "//app/interface/passport-register/model:go_default_library",
        "//app/interface/passport-register/service:go_default_library",
        "//library/ecode:go_default_library",
        "//library/log:go_default_library",
        "//library/net/http/mars:go_default_library",
    ],
)
