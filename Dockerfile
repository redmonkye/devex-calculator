# ===== Build stage =====
# Fixed: Updated to 1.26 to satisfy go.mod requirement (>= 1.25.6)
FROM golang:1.25-alpine AS builder
WORKDIR /app

# Copy dependency files first for better caching
COPY go.mod ./
# COPY go.sum ./  <-- Uncomment if you have this file

# Copy the rest of the application
COPY . .

# Build the binary
RUN go build -o main .

# ===== Runtime stage =====
FROM alpine:latest
WORKDIR /root/

# Fixed: Install certificates in the FINAL stage, not a temporary one
RUN apk add --no-cache ca-certificates

# Copy the compiled binary
COPY --from=builder /app/main .

# Fixed: Explicitly copy the templates folder so the app can find index.html
COPY --from=builder /app/templates ./templates

EXPOSE 8080
CMD ["./main"]