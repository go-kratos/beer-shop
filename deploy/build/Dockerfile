FROM golang:1.15 AS builder

ARG APP_NAME

COPY . /src
WORKDIR /src/app/${APP_NAME}

RUN mkdir -p /release && GOPROXY=https://goproxy.cn go build -o /release ./...

FROM debian:stable-slim

RUN apt-get update && apt-get install -y --no-install-recommends \
		ca-certificates  \
        netbase \
        && rm -rf /var/lib/apt/lists/ \
        && apt-get autoremove -y && apt-get autoclean -y

COPY --from=builder /release /app

WORKDIR /app

EXPOSE 8000
EXPOSE 9000
VOLUME /data/conf

CMD ["./server", "-conf", "/data/conf"]