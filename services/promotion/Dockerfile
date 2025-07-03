# Build stage
FROM golang:1.24-alpine AS builder

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o promotion ./services/promotion/cmd/api

# Final stage
FROM alpine:latest AS final

COPY --from=builder /build/promotion /promotion

EXPOSE 8081
EXPOSE 50051

RUN chmod +x /promotion
ENTRYPOINT ["/promotion"]
