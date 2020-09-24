ARG arch=amd64

MAINTAINER "GitHub shokujinjp Team <https://github.com/shokujinjp>"

FROM ${arch}/alpine

RUN apk add --no-cache ca-certificates
ADD ./api /api

EXPOSE 8080
ENTRYPOINT ["/api"]