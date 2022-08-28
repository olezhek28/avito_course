FROM golang:1.18.5-alpine3.16 AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED 0
ENV GOOS linux

RUN apk update --no-cache && apk add --no-cache tzdata
RUN apk add make

WORKDIR /build

COPY . .

RUN make build-scheduler

FROM alpine:3.16

LABEL maintainer=godev@okozyrev.ru

RUN apk update --no-cache && apk add --no-cache ca-certificates

WORKDIR /bin
COPY --from=builder build/bin/scheduler .
COPY --from=builder build/configs/scheduler_config.json .
CMD ["./scheduler"]