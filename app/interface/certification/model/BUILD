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
        "arg_work_cert.go",
        "const.go",
        "model.go",
        "resp_id_cert.go",
        "resp_work_cert.go",
    ],
    importpath = "valerian/app/interface/certification/model",
    deps = ["//library/database/sqlx/types:go_default_library"],
)
