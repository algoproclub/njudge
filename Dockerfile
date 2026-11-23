FROM golang:1.27-trixie

WORKDIR /app

ARG DEBIAN_FRONTEND=noninteractive
ENV TZ=Europe/Budapest
RUN apt-get update && apt-get install -y --no-install-recommends pandoc \
    && rm -rf /var/lib/apt/lists/*

COPY go.mod go.sum ./
RUN go mod download && go install github.com/go-delve/delve/cmd/dlv@latest

COPY cmd/ ./cmd
COPY pkg/ ./pkg
COPY internal/ ./internal
COPY main.go ./

RUN go build
