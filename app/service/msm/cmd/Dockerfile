FROM golang:1.12.7-alpine as builder
RUN apk update && apk add --no-cache git ca-certificates tzdata && update-ca-certificates
RUN adduser -D -g '' appuser

# WORKDIR $GOPATH/src/valerian/app/infra/discovery/cmd
# RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -a -installsuffix cgo -o /go/bin/cmd .

FROM alpine:3.10.2
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
RUN  cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && echo "Asia/Shanghai" > /etc/timezone
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd
COPY ./cmd /go/bin/cmd
USER appuser

# Run the hello binary.
CMD ["/go/bin/cmd", "--conf","/go/bin/config/config.toml"]

