FROM golang:1.16.6-alpine3.14 as builder
ENV CGO_ENABLED=0
ADD . /go/src/twemproxy-exporter
WORKDIR /go/src/twemproxy-exporter
RUN go get
RUN go build && chmod +x twemproxy-exporter

FROM debian:stretch-slim
ARG TWEMPROXY_TARGET_HOST
ARG TWEMPROXY_TARGET_PORT
ARG TWEMPROXY_EXPORTER_PORT
COPY --from=builder /go/src/twemproxy-exporter/twemproxy-exporter /bin/twemproxy-exporter
ENTRYPOINT ["/bin/twemproxy-exporter"]
