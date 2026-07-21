FROM golang:1.25.4-alpine AS builder
RUN apk add --no-cache git
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o user-service ./cmd/server/main.go
FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /app/user-service /user-service
COPY --from=builder /app/db/migrations /db/migrations
EXPOSE 50052
CMD ["/user-service"]