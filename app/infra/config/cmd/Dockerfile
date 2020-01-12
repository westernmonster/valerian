FROM registry.cn-hangzhou.aliyuncs.com/valerian/base:latest
COPY config.toml /go/bin/config.toml
COPY cmd /go/bin/cmd


# Run the hello binary.
CMD ["/go/bin/cmd","--conf=/go/bin/config.toml"]


