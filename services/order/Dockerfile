# Build stage
FROM golang:1.24-alpine AS builder

WORKDIR /build

COPY go.mod go.sum ./
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -mod=vendor -o order ./services/order/cmd/api

# Final stage
FROM alpine:latest AS final

COPY --from=builder /build/order /order

EXPOSE 8080
EXPOSE 50050

RUN chmod +x /order
ENTRYPOINT ["/order"]
