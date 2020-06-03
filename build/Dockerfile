FROM golang:1.12-alpine

WORKDIR /go/src/github.com/sillent/redis-failover
COPY . .
RUN build/build.sh

FROM alpine:latest
RUN apk --no-cache add \
  ca-certificates
COPY --from=0 /go/src/github.com/sillent/redis-failover/bin/linux/failover /usr/local/bin
RUN addgroup -g 1000 rf && \
    adduser -D -u 1000 -G rf rf && \
    chown rf:rf /usr/local/bin/failover
USER rf

ENTRYPOINT ["/usr/local/bin/failover"]

