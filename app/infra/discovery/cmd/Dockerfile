FROM golang:1.12.7-alpine as builder
RUN apk update && apk add --no-cache git ca-certificates tzdata && update-ca-certificates
RUN adduser -D -g '' appuser

# WORKDIR $GOPATH/src/valerian/app/infra/discovery/cmd
# RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -a -installsuffix cgo -o /go/bin/cmd .

FROM scratch
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd
# COPY ./discovery-idc1-1.toml /go/bin/discovery-idc1-1.toml
COPY ./cmd /go/bin/cmd
USER appuser

# Run the hello binary.
ENTRYPOINT ["/go/bin/cmd", "-conf","/go/bin/config.toml"]
