FROM golang:1.15.6-alpine3.12 AS builder
RUN apk update
RUN apk -U --no-cache add libpcap-dev build-base
WORKDIR /app
COPY printproxy.go go.mod ./
RUN go mod download
RUN go build -ldflags="-extldflags=-static" -o proxy .

FROM alpine:latest
COPY --from=builder /app/proxy .
RUN chmod +x proxy
CMD ["./proxy"]