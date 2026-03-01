# Multi-stage build (single image) that builds the binary and produces a runnable image
FROM golang:1.26 AS builder

WORKDIR /src

# Provide build dependencies for cgo (sqlite3)
RUN apt-get update && \
    apt-get install -y --no-install-recommends build-essential gcc ca-certificates && \
    rm -rf /var/lib/apt/lists/*

# Cache modules
COPY go.mod go.sum ./
RUN go mod download || true

# Copy sources and build
COPY . .
RUN CGO_ENABLED=1 go build -ldflags "-s -w" -o /usr/local/bin/go-via .

# Final image
FROM debian:bookworm-slim

# required runtime CA certs and lib for sqlite
RUN apt-get update && \
    apt-get install -y --no-install-recommends ca-certificates libsqlite3-0 && \
    rm -rf /var/lib/apt/lists/*

COPY --from=builder /usr/local/bin/go-via /usr/local/bin/go-via

EXPOSE 8080
ENTRYPOINT ["/usr/local/bin/go-via"]
