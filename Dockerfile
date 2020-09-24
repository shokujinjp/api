ARG arch=amd64

FROM ${arch}/alpine

MAINTAINER "GitHub shokujinjp Team <https://github.com/shokujinjp>"

RUN apk add --no-cache ca-certificates
ADD ./api /api

EXPOSE 8080
ENTRYPOINT ["/api"]