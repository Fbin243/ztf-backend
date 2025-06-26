# Build stage
FROM golang:1.24-alpine AS builder

WORKDIR /build

COPY go.mod go.sum ./
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -mod=vendor -o promotion-api ./promotion/cmd/api

# Final stage
FROM alpine:latest AS final

COPY --from=builder /build/promotion-api /promotion-api

RUN chmod +x /promotion-api

EXPOSE 8081
EXPOSE 50051

ENTRYPOINT ["/promotion-api"]
