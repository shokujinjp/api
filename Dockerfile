FROM golang:latest as builder

MAINTAINER "GitHub shokujinjp Team <https://github.com/shokujinjp>"

ENV CGO_ENABLED=0
WORKDIR /go/src/github.com/shokujinjp/api
COPY . .

ENV GO111MODULE=on

RUN go mod download
RUN go build .

# runtime image
FROM alpine
RUN apk add --no-cache ca-certificates
COPY --from=builder /go/src/github.com/shokujinjp/api/api /api
EXPOSE 8080
ENTRYPOINT ["/api"]