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
        "//vendor/go.uber.org/zap/buffer:all-srcs",
        "//vendor/go.uber.org/zap/internal/bufferpool:all-srcs",
        "//vendor/go.uber.org/zap/internal/color:all-srcs",
        "//vendor/go.uber.org/zap/internal/exit:all-srcs",
        "//vendor/go.uber.org/zap/zapcore:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
