FROM golang@sha256:cee6f4b901543e8e3f20da3a4f7caac6ea643fd5a46201c3c2387183a332d989 as builder
RUN apk update && apk add --no-cache git ca-certificates tzdata && update-ca-certificates
RUN adduser -D -g '' appuser
FROM scratch
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd
COPY ./cmd /go/bin/cmd
COPY ./discovery-idc1-1.toml /go/bin/discovery-idc1-1.toml
USER appuser

# Run the hello binary.
ENTRYPOINT ["/go/bin/cmd", "-conf","/go/bin/discovery-idc1-1.toml"]
