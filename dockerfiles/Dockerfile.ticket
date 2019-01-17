FROM golang:alpine AS build-env
ADD . /go/src/hex-example
RUN apk update \
    && apk upgrade \
    && apk add git \
    && cd /go/src/hex-example \
    && go get ./... \
    && CGO_ENABLED=0 GOOS=linux go build -ldflags "-w -X main.docker=true" cmd/ticketAPI/main.go

FROM scratch
LABEL maintainer="Joel Holmes <holmes89@gmail.com>"
ENV PORT 3000
EXPOSE 3000
ENV DATABASE_URL=""
ENV REDIS_PASSWORD=""
COPY --from=build-env /go/src/hex-example/main /
CMD ["/main"]
