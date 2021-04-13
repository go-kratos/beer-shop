FROM golang:1.15 AS builder

COPY ${PWD} /src
WORKDIR /src

RUN mkdir -p bin/ && go build -ldflags "-X main.Version=$(shell git describe --tags --always)" -o ./bin/ ./...

FROM debian:stable-slim

RUN apt-get update && apt-get install -y --no-install-recommends \
		ca-certificates  \
        netbase \
        && rm -rf /var/lib/apt/lists/ \
        && apt-get autoremove -y && apt-get autoclean -y

COPY --from=builder /src/bin /app

WORKDIR /app

EXPOSE 8000
EXPOSE 9000

CMD ["./user"]