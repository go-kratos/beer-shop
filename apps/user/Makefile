GOPATH:=$(shell go env GOPATH)
VERSION=$(shell git describe --tags --always)
PROTO_FILES=$(shell find . -name *.proto)
KRATOS_VERSION=$(shell go mod graph |grep go-kratos/kratos/v2 |head -n 1 |awk -F '@' '{print $$2}')
KRATOS=$(GOPATH)/pkg/mod/github.com/go-kratos/kratos/v2@$(KRATOS_VERSION)

.PHONY: init
# init env
init:
	go get -u google.golang.org/protobuf/cmd/protoc-gen-go
	go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc
	go get -u github.com/go-kratos/kratos/cmd/protoc-gen-go-http/v2
	go get -u github.com/go-kratos/kratos/cmd/protoc-gen-go-errors/v2
	go get -u github.com/google/wire/cmd/wire

.PHONY: grpc
# generate grpc code
grpc:
	protoc --proto_path=. \
           --proto_path=$(KRATOS)/api \
           --proto_path=$(KRATOS)/third_party \
           --proto_path=$(GOPATH)/src \
           --go_out=paths=source_relative:. \
           --go-grpc_out=paths=source_relative:. \
           --go-errors_out=paths=source_relative:. \
           $(PROTO_FILES)

.PHONY: http
# generate http code
http:
	protoc --proto_path=. \
           --proto_path=$(KRATOS)/api \
           --proto_path=$(KRATOS)/third_party \
           --proto_path=$(GOPATH)/src \
           --go_out=paths=source_relative:. \
           --go-http_out=paths=source_relative:. \
           --go-errors_out=paths=source_relative:. \
           $(PROTO_FILES)

.PHONY: generate
# generate client code
generate:
	go generate ./...

.PHONY: build
# build
build:
	mkdir -p bin/ && go build -ldflags "-X main.Version=$(VERSION)" -o ./bin/ ./...

.PHONY: test
# test
test:
	go test -v ./... -cover

.PHONY: run
run:
	cd cmd/user/ && go run .

.PHONY: ent
ent:
	cd internal/data/ && ent generate ./ent/schema

.PHONY: all
# generate all
all:
	make grpc;
	make http;
	make generate;
	make build;
	make test;

# show help
help:
	@echo ''
	@echo 'Usage:'
	@echo ' make [target]'
	@echo ''
	@echo 'Targets:'
	@awk '/^[a-zA-Z\-\_0-9]+:/ { \
	helpMessage = match(lastLine, /^# (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")-1); \
			helpMessage = substr(lastLine, RSTART + 2, RLENGTH); \
			printf "\033[36m%-22s\033[0m %s\n", helpCommand,helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help
