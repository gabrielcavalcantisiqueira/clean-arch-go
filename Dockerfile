FROM golang:1.23 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o clean-arch ./cmd/ordersystem/.

FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /app/clean-arch /
COPY ./cmd/ordersystem/.env /
EXPOSE 8000 50051 8080
CMD ["/clean-arch"]
