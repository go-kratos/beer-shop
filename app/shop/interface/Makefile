include ../../../app_makefile

API_PROTO_FILES=$(shell cd ../../../ && find api -name '*.proto'|awk '{print "../../../"$$1}')

.PHONY: api
# generate api code
api:
	 protoc --proto_path=../../../api \
           --proto_path=../../../third_party \
 	       --go_out=paths=import,module=$(MODULE):$(API_PROTOC_OUTPUT_PATH) \
 	       --go-http_out=paths=import,module=$(MODULE):$(API_PROTOC_OUTPUT_PATH) \
 	       --go-grpc_out=paths=import,module=$(MODULE):$(API_PROTOC_OUTPUT_PATH) \
           $(API_PROTO_FILES)