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
        "arg.go",
        "arg_block.go",
        "arg_follow.go",
        "arg_setting.go",
        "const.go",
        "model.go",
        "resp.go",
        "resp_follow.go",
        "resp_member_article.go",
        "resp_member_discuss.go",
        "resp_member_feed.go",
        "resp_member_info.go",
        "resp_member_topic.go",
        "resp_setting.go",
    ],
    importpath = "valerian/app/interface/account/model",
    deps = [
        "//library/database/sqlx/types:go_default_library",
        "//library/ecode:go_default_library",
        "//vendor/github.com/asaskevich/govalidator:go_default_library",
        "//vendor/github.com/go-ozzo/ozzo-validation:go_default_library",
        "//vendor/github.com/go-ozzo/ozzo-validation/is:go_default_library",
    ],
)
