FROM registry.cn-hangzhou.aliyuncs.com/valerian/base:latest
ARG EXE
COPY ${EXE} /go/bin/cmd
RUN touch /go/bin/config.toml


# Run the hello binary.
CMD ["/go/bin/cmd"]


