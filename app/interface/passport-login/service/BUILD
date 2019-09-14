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
        "generate.go",
        "identity.go",
        "login.go",
        "md5.go",
        "oauth.go",
        "service.go",
        "util.go",
    ],
    importpath = "valerian/app/interface/passport-login/service",
    deps = [
        "//app/interface/passport-login/conf:go_default_library",
        "//app/interface/passport-login/dao:go_default_library",
        "//app/interface/passport-login/model:go_default_library",
        "//library/database/sqalx:go_default_library",
        "//library/ecode:go_default_library",
        "//library/gid:go_default_library",
        "//library/log:go_default_library",
        "//vendor/github.com/satori/go.uuid:go_default_library",
    ],
)
