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
        "account_topic_setting.go",
        "dao.go",
        "discuss_category.go",
        "mc.go",
        "mc_account_topic_setting.go",
        "mc_topic.go",
        "mc_topic_auth.go",
        "mc_topic_catalog.go",
        "mc_topic_member.go",
        "topic.go",
        "topic_auth.go",
        "topic_catalog.go",
        "topic_follow.go",
        "topic_invite_request.go",
        "topic_member.go",
        "topic_member_stat.go",
    ],
    importpath = "valerian/app/service/topic/dao",
    deps = [
        "//app/service/account/api:go_default_library",
        "//app/service/topic/conf:go_default_library",
        "//app/service/topic/model:go_default_library",
        "//library/cache/memcache:go_default_library",
        "//library/database/sqalx:go_default_library",
        "//library/log:go_default_library",
        "//library/stat/prom:go_default_library",
        "//vendor/github.com/satori/go.uuid:go_default_library",
    ],
)
