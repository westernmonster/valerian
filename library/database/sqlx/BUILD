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
        "//library/database/sqlx/reflectx:all-srcs",
        "//library/database/sqlx/types:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)

go_library(
    name = "go_default_library",
    srcs = [
        "bind.go",
        "doc.go",
        "sqlx.go",
        "sqlx_context.go",
    ],
    importpath = "valerian/library/database/sqlx",
    deps = [
        "//library/database/sqlx/reflectx:go_default_library",
        "//library/net/netutil/breaker:go_default_library",
        "//library/stat:go_default_library",
        "//library/time:go_default_library",
        "//library/tracing:go_default_library",
        "//vendor/github.com/jmoiron/sqlx/reflectx:go_default_library",
        "//vendor/github.com/opentracing/opentracing-go:go_default_library",
        "//vendor/github.com/pkg/errors:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = [
        "sqlx_context_test.go",
        "sqlx_test.go",
    ],
    embed = [":go_default_library"],
    deps = [
        "//vendor/github.com/go-sql-driver/mysql:go_default_library",
        "//vendor/github.com/jmoiron/sqlx/reflectx:go_default_library",
        "//vendor/github.com/lib/pq:go_default_library",
        "//vendor/github.com/mattn/go-sqlite3:go_default_library",
    ],
)
