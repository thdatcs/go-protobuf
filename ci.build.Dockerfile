ARG IMAGE

FROM base-${IMAGE} as builder
LABEL stage=${IMAGE}
RUN make build && \
    mkdir build && \
    mkdir -p build/statik && \
    cp -p $GOPATH/src/go-protobuf/cmd/go-protobuf $GOPATH/build && \
    cp -p $GOPATH/src/go-protobuf/cmd/config.yaml $GOPATH/build && \
    cp -p $GOPATH/src/go-protobuf/cmd/statik/statik.go $GOPATH/build/statik

FROM alpine:latest
RUN mkdir -p /app
WORKDIR /app
RUN apk add --no-cache ca-certificates openssl
COPY --from=builder /go/build /app
CMD ["./go-protobuf", "-configFile=config.yaml"]
