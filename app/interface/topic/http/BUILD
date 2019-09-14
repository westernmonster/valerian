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
        "article.go",
        "article_file.go",
        "article_history.go",
        "article_relation.go",
        "auth_topic.go",
        "comment.go",
        "discuss_categories.go",
        "http.go",
        "invite.go",
        "revise.go",
        "search.go",
        "topic.go",
        "topic_activity.go",
        "topic_catalog.go",
        "topic_discuss.go",
        "topic_follow.go",
        "topic_member.go",
    ],
    importpath = "valerian/app/interface/topic/http",
    deps = [
        "//app/interface/topic/conf:go_default_library",
        "//app/interface/topic/model:go_default_library",
        "//app/interface/topic/service:go_default_library",
        "//library/ecode:go_default_library",
        "//library/log:go_default_library",
        "//library/net/http/mars:go_default_library",
        "//library/net/http/mars/middleware/auth:go_default_library",
    ],
)
