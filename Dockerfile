FROM golang:1.24.2 AS base

ENV CGO_ENABLED=0 GOOS=linux

WORKDIR /app

FROM base AS deps

COPY go.mod go.sum ./
RUN go mod download

FROM base AS builder

COPY . .
RUN go build -o wbmagalhaes .

FROM scratch

WORKDIR /

COPY --from=builder /app/wbmagalhaes /wbmagalhaes

EXPOSE 3000

ENTRYPOINT ["/wbmagalhaes"]
