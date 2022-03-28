# syntax=docker:dockerfile/1

FROM golang:alpine

RUN mkdir -p /src/build
WORKDIR  /src/build

RUN apk add --no-cache tzdata ca-certificates

COPY main ./main

EXPOSE 8000

CMD ["./main"]
