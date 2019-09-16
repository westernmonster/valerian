FROM registry.cn-hangzhou.aliyuncs.com/valerian/base:latest
COPY $EXE /go/bin/cmd
USER appuser

# Run the hello binary.
CMD ["/go/bin/cmd", "--conf","/go/bin/config/config.toml"]


