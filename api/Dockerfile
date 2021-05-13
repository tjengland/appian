FROM golang:alpine3.13 as base

ADD . $GOPATH/src/github.com/tjengland/appian/api
WORKDIR $GOPATH/src/github.com/tjengland/appian/api
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=arm64
ENV GO111MODULE=off
RUN go build -o appian-api
RUN chmod +x appian-api
ENV PORT="8081"
EXPOSE ${PORT}
