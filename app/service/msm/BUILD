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
        "//app/service/msm/cmd:all-srcs",
        "//app/service/msm/conf:all-srcs",
        "//app/service/msm/dao:all-srcs",
        "//app/service/msm/http:all-srcs",
        "//app/service/msm/model:all-srcs",
        "//app/service/msm/service:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
