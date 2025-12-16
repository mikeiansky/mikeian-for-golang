FROM golang:1.25 AS builder

# 安装 golangci-lint（官方推荐方式）
RUN curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | \
    sh -s -- -b $(go env GOPATH)/bin v1.55.2

# 确保 go bin 目录在 PATH 中
ENV PATH="$PATH:$(go env GOPATH)/bin"

COPY . /src
WORKDIR /src

RUN make hello