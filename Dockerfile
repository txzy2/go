FROM golang:1.24-bookworm AS builder

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

RUN go build -v -o /app/first ./cmd/first/
FROM debian:bookworm-slim

WORKDIR /app
COPY --from=builder /app/first /app/first

RUN chmod +x /app/first

ENTRYPOINT ["/app/first"]
