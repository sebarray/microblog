FROM golang:latest AS builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64
WORKDIR /go/src
COPY go.mod .
RUN go mod download
COPY . .
RUN go build ./cmd/query_service/main.go

FROM alpine
COPY --from=builder /go/src .
ENTRYPOINT  ["./main"]