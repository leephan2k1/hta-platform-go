# ---- Stage 1: Build ----
FROM golang:1.26-alpine AS builder

# Install git & ca-certificates (needed for fetching Go modules over HTTPS)
RUN apk add --no-cache git ca-certificates tzdata

WORKDIR /app

# Cache dependencies first – these layers are invalidated only when go.mod/go.sum change
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build a fully static binary (no CGO) for the smallest possible image
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -ldflags="-s -w" -o /app/server ./cmd/server

# ---- Stage 2: Runtime ----
FROM alpine:3.22

# Install only the bare minimum for a production container
RUN apk add --no-cache ca-certificates tzdata

# Non-root user for security
RUN addgroup -S appgroup && adduser -S appuser -G appgroup

WORKDIR /app

# Copy the compiled binary from the builder stage
COPY --from=builder /app/server .

# Own the workdir
RUN chown -R appuser:appgroup /app

USER appuser

EXPOSE 8800

ENTRYPOINT ["./server"]
