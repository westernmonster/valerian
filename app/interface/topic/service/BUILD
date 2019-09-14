package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_test",
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
        "//app/interface/topic/service/mocks:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)

go_library(
    name = "go_default_library",
    srcs = [
        "account.go",
        "article.go",
        "article_file.go",
        "article_history.go",
        "article_relation.go",
        "auth_topic.go",
        "discuss_category.go",
        "interface.go",
        "service.go",
        "topic.go",
        "topic_catalog.go",
        "topic_follow.go",
        "topic_invite.go",
        "topic_member.go",
        "topic_meta.go",
        "topic_search.go",
    ],
    importpath = "valerian/app/interface/topic/service",
    deps = [
        "//app/interface/topic/conf:go_default_library",
        "//app/interface/topic/dao:go_default_library",
        "//app/interface/topic/model:go_default_library",
        "//library/conf/env:go_default_library",
        "//library/database/sqalx:go_default_library",
        "//library/database/sqlx/types:go_default_library",
        "//library/ecode:go_default_library",
        "//library/gid:go_default_library",
        "//library/log:go_default_library",
        "//library/net/metadata:go_default_library",
        "//vendor/github.com/jinzhu/copier:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = ["topic_member_test.go"],
    embed = [":go_default_library"],
    deps = [
        "//app/interface/topic/model:go_default_library",
        "//app/interface/topic/service/mocks:go_default_library",
        "//library/net/metadata:go_default_library",
        "//vendor/github.com/stretchr/testify/assert:go_default_library",
        "//vendor/github.com/stretchr/testify/mock:go_default_library",
    ],
)
