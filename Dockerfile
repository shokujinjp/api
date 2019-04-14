FROM golang:latest as builder

MAINTAINER "GitHub shokujinjp Team <https://github.com/shokujinjp>"

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
WORKDIR /go/src/github.com/shokujinjp/api
COPY . .
RUN go get -u github.com/gorilla/handlers \
    && go get -u github.com/gorilla/mux \
    && go get -u github.com/shokujinjp/shokujinjp-sdk-go/shokujinjp
RUN go build . 

# runtime image
FROM alpine
RUN apk add --no-cache ca-certificates
COPY --from=builder /go/src/github.com/shokujinjp/api/api /api
EXPOSE 8080
ENTRYPOINT ["/api"]
