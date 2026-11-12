# Build stage
ARG GO_VERSION=alpine
FROM golang:${GO_VERSION} AS builder

WORKDIR /app

# Install templ
RUN go install github.com/a-h/templ/cmd/templ@latest

# Dependencies
COPY go.mod go.sum ./
RUN go mod download

# Source code
COPY . .

# Generate templ files
RUN templ generate

# Build binary
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /server ./cmd/server

# Runtime stage
FROM alpine:3.20

WORKDIR /app

# Run as non-root user
RUN adduser -D -u 10001 rhyred

COPY --from=builder --chown=rhyred:rhyred /server .
COPY --from=builder --chown=rhyred:rhyred /app/web/static ./web/static
COPY --from=builder --chown=rhyred:rhyred /app/web/data ./web/data

USER rhyred:rhyred

EXPOSE 8080

HEALTHCHECK --interval=30s --timeout=3s --retries=3 \
  CMD wget --no-verbose --tries=1 --spider http://localhost:8080/health || exit 1

ENTRYPOINT ["./server"]
