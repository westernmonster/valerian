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
        "account_article_attr.go",
        "account_follow.go",
        "account_search.go",
        "account_topic_setting.go",
        "article.go",
        "article_file.go",
        "article_history.go",
        "auth_topic.go",
        "dao.go",
        "discuss_category.go",
        "es.go",
        "locale.go",
        "mc.go",
        "mc_account.go",
        "mc_account_topic_setting.go",
        "mc_article.go",
        "mc_article_file.go",
        "mc_article_history.go",
        "mc_auth_topic.go",
        "mc_topic.go",
        "mc_topic_catalog.go",
        "mc_topic_member.go",
        "topic.go",
        "topic_catalog.go",
        "topic_follow.go",
        "topic_invite_request.go",
        "topic_member.go",
    ],
    importpath = "valerian/app/interface/topic/dao",
    deps = [
        "//app/interface/topic/conf:go_default_library",
        "//app/interface/topic/model:go_default_library",
        "//library/cache/memcache:go_default_library",
        "//library/database/sqalx:go_default_library",
        "//library/log:go_default_library",
        "//library/stat/prom:go_default_library",
        "//vendor/github.com/satori/go.uuid:go_default_library",
        "//vendor/gopkg.in/olivere/elastic.v6:go_default_library",
    ],
)
