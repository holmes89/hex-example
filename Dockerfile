FROM golang:alpine AS build-env
ADD . /go/src/hex-example
RUN apk update \
    && apk upgrade \
    && apk add git \
    && cd /go/src/hex-example \
    && go get ./... \
    && CGO_ENABLED=0 GOOS=linux go build main.go

FROM scratch
MAINTAINER "Joel Holmes <holmes89@gmail.com>"
ENV PORT 3000
EXPOSE 3000
COPY --from=build-env /go/src/hex-example /
CMD ["/main"]
