FROM golang:1.21.5-alpine AS builder

RUN apk --update add ca-certificates git

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o main cmd/server/main.go

# Build a smaller image that will only contain the application's binary
FROM scratch


# Move to working directory /app
WORKDIR /app


# Copy certificates
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt


# Copy application's binary
COPY --from=builder /app .

EXPOSE 9200

# Command to run the application when starting the container
CMD ["./main"]