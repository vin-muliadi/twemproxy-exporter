# Note you cannot run golang binaries on Alpine directly
FROM            debian:buster-slim

MAINTAINER      chris.mague@shokunin.co

COPY            twemproxy-exporter /twemproxy-exporter

WORKDIR		/
ENV		GIN_MODE=release

EXPOSE          8080

ENTRYPOINT      [ "/twemproxy-exporter" ]
