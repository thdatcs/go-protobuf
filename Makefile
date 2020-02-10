include ci.mk

GOPATH := $(PWD)
GOBIN := $(GOPATH)/bin
PATH := "$(PATH):$(GOPATH)/bin"
PKGS := go-protobuf
TESTPKGS := go-protobuf/servers

.PHONY:

prepare:
	GOPATH=$(GOPATH) GOBIN=$(GOBIN) go get -v -u github.com/golang/dep/cmd/dep
	GOPATH=$(GOPATH) GOBIN=$(GOBIN) go get -v -u github.com/rakyll/statik
	GOPATH=$(GOPATH) GOBIN=$(GOBIN) go get -v -u golang.org/x/lint/golint
	GOPATH=$(GOPATH) GOBIN=$(GOBIN) go get -v -u github.com/golangci/golangci-lint/cmd/golangci-lint
	GOPATH=$(GOPATH) GOBIN=$(GOBIN) go get -v -u github.com/vektra/mockery/.../
	GOPATH=$(GOPATH) GOBIN=$(GOBIN) go get -v -u github.com/golang/protobuf/protoc-gen-go
	GOPATH=$(GOPATH) GOBIN=$(GOBIN) go get -v -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
	GOPATH=$(GOPATH) GOBIN=$(GOBIN) go get -v -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
	GOPATH=$(GOPATH) GOBIN=$(GOBIN) go get -v -u github.com/favadi/protoc-go-inject-tag

dep:
	bash $(GOPATH)/scripts/dep.sh $(PKGS)

mock:
	GOPATH=${GOPATH} ${GOPATH}/bin/mockery -name $(interface) -dir $(directory) -output $(directory)/mocks

lint:
	bash $(GOPATH)/scripts/golangci-lint.sh $(PKGS)
	bash $(GOPATH)/scripts/golint.sh $(PKGS)

unit-test:
	GOPATH=${GOPATH} go test ${TESTPKGS} -v
	GOPATH=${GOPATH} go test -cover ${TESTPKGS}

integration-test:
	cd $(GOPATH)/src/go-protobuf/integrations && newman run go-protobuf.postman_collection.json -e go-protobuf.postman_environment.json

build:
	cd $(GOPATH)/src/go-protobuf/cmd && CGO_ENABLED=0 GOOS=linux GOPATH=$(GOPATH) go build -a -installsuffix cgo -o go-protobuf .

run:
	cd $(GOPATH)/src/go-protobuf/cmd && GOPATH=$(GOPATH) go run main.go -configFile=config.yaml
