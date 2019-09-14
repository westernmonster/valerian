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
        "account.go",
        "follow.go",
        "http.go",
        "member.go",
        "setting.go",
    ],
    importpath = "valerian/app/interface/account/http",
    deps = [
        "//app/interface/account/conf:go_default_library",
        "//app/interface/account/model:go_default_library",
        "//app/interface/account/service:go_default_library",
        "//library/ecode:go_default_library",
        "//library/log:go_default_library",
        "//library/net/http/mars:go_default_library",
        "//library/net/http/mars/middleware/auth:go_default_library",
    ],
)
