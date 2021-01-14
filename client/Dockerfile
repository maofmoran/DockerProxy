FROM golang:1.14 AS builder
ENV CGO_ENABLED=0
WORKDIR /app
COPY client.go ./
RUN go build -o client .

FROM alpine:latest
COPY --from=builder /app/client .
RUN chmod +x client
CMD ["./client"]