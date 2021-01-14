FROM golang:1.14 AS builder
ENV CGO_ENABLED=0
WORKDIR /app
COPY server.go go.mod go.sum ./
RUN go mod download
RUN go build -o server .

FROM alpine:latest
COPY --from=builder /app/server .
RUN chmod +x server
CMD ["./server"]