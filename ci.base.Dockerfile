ARG IMAGE

FROM golang
ADD . /go
WORKDIR /go
RUN make prepare && \
    make dep