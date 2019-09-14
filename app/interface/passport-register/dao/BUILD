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
        "area.go",
        "auth_mc.go",
        "dao.go",
        "mc.go",
        "token.go",
    ],
    importpath = "valerian/app/interface/passport-register/dao",
    deps = [
        "//app/interface/passport-register/conf:go_default_library",
        "//app/interface/passport-register/model:go_default_library",
        "//library/cache/memcache:go_default_library",
        "//library/database/sqalx:go_default_library",
        "//library/log:go_default_library",
    ],
)
