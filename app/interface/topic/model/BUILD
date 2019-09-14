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
        "//app/interface/topic/model/docs:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)

go_library(
    name = "go_default_library",
    srcs = [
        "arg_article.go",
        "arg_article_file.go",
        "arg_article_relation.go",
        "arg_auth_topic.go",
        "arg_comment.go",
        "arg_common.go",
        "arg_revise.go",
        "arg_topic.go",
        "arg_topic_catalog.go",
        "arg_topic_discuss.go",
        "arg_topic_disucss_category.go",
        "arg_topic_follow.go",
        "arg_topic_member.go",
        "const.go",
        "es.go",
        "model.go",
        "model_account.go",
        "model_account_follow.go",
        "model_article.go",
        "model_locale.go",
        "resp_account.go",
        "resp_article.go",
        "resp_article_history.go",
        "resp_article_relation.go",
        "resp_auth_topic.go",
        "resp_comment.go",
        "resp_discuss.go",
        "resp_discuss_category.go",
        "resp_feed.go",
        "resp_member_fans.go",
        "resp_revise.go",
        "resp_topic.go",
        "resp_topic_catalog.go",
        "resp_topic_member.go",
        "resp_topic_search.go",
    ],
    importpath = "valerian/app/interface/topic/model",
    deps = [
        "//library/database/sqlx/types:go_default_library",
        "//library/ecode:go_default_library",
        "//vendor/github.com/go-ozzo/ozzo-validation:go_default_library",
        "//vendor/github.com/go-ozzo/ozzo-validation/is:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = ["arg_topic_catalog_test.go"],
    embed = [":go_default_library"],
    deps = ["//vendor/github.com/stretchr/testify/assert:go_default_library"],
)
