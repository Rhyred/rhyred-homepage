# Build stage
FROM golang:1.27-alpine AS builder

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

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /server ./cmd/server

# Runtime stage
FROM alpine:3.19

WORKDIR /app

COPY --from=builder /server .
COPY --from=builder /app/web/static ./web/static
COPY --from=builder /app/web/data ./web/data

EXPOSE 8080

ENTRYPOINT ["./server"]

