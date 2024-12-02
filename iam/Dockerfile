# syntax=docker/dockerfile:1.2
FROM alpine:3.20

ARG TARGETPLATFORM
COPY ./dist/$TARGETPLATFORM/zziam /

EXPOSE 8080

ENTRYPOINT ["/zziam"]
