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
    srcs = [
        ":package-srcs",
        "//vendor/gopkg.in/h2non/gentleman.v2/context:all-srcs",
        "//vendor/gopkg.in/h2non/gentleman.v2/middleware:all-srcs",
        "//vendor/gopkg.in/h2non/gentleman.v2/mux:all-srcs",
        "//vendor/gopkg.in/h2non/gentleman.v2/plugin:all-srcs",
        "//vendor/gopkg.in/h2non/gentleman.v2/plugins/body:all-srcs",
        "//vendor/gopkg.in/h2non/gentleman.v2/plugins/bodytype:all-srcs",
        "//vendor/gopkg.in/h2non/gentleman.v2/plugins/cookies:all-srcs",
        "//vendor/gopkg.in/h2non/gentleman.v2/plugins/headers:all-srcs",
        "//vendor/gopkg.in/h2non/gentleman.v2/plugins/multipart:all-srcs",
        "//vendor/gopkg.in/h2non/gentleman.v2/plugins/query:all-srcs",
        "//vendor/gopkg.in/h2non/gentleman.v2/plugins/url:all-srcs",
        "//vendor/gopkg.in/h2non/gentleman.v2/utils:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
