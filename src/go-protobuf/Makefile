GOPATH := $(PWD)/../..
GOBIN := $(GOPATH)/bin
PATH := "$(PATH):$(GOPATH)/bin"
PROJECT_NAME := "idm-system"

# make gen-swagger
gen-swagger:
	@PATH=$(PATH) $(GOPATH)/third_party/protoc \
		--proto_path=proto \
		--proto_path=$(GOPATH)/third_party/google \
		--swagger_out=allow_delete_body=true,allow_merge=true,logtostderr=true:swagger \
		api_user.proto api_ping.proto api_swagger.proto
	@PATH=$(PATH) statik -src=swagger -dest=cmd

# make proto-file=... gen-api
gen-api: gen-grpc gen-grpc-gateway

gen-grpc:
	@PATH=$(PATH) $(GOPATH)/third_party/protoc \
		--proto_path=proto \
		--proto_path=$(GOPATH)/third_party/google \
		--go_out=plugins=grpc:api \
		$(proto-file)

gen-grpc-gateway:
	@PATH=$(PATH) $(GOPATH)/third_party/protoc \
		--proto_path=proto \
		--proto_path=$(GOPATH)/third_party/google \
		--grpc-gateway_out=logtostderr=true:api \
		$(proto-file)