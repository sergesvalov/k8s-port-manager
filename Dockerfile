# Build stage
FROM golang:1.21-alpine AS builder
WORKDIR /app

# Copy dependency files first
COPY go.mod go.sum ./
RUN go mod download

# Copy source code and templates
COPY main.go .
COPY templates/ ./templates/

# Build
RUN go build -o port-manager main.go

# Run stage
FROM alpine:3.18
WORKDIR /app

# Copy binary and templates from builder
COPY --from=builder /app/port-manager .
COPY --from=builder /app/templates ./templates

# Create data directory
RUN mkdir /data 
EXPOSE 8080
CMD ["./port-manager"]