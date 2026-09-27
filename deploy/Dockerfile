FROM golang:1.27.1-alpine AS builder

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Application
RUN CGO_ENABLED=0 go build \
    -trimpath \
    -ldflags="-s -w" \
    -o /out/simpleclansstats \
    ./cmd/server

# Goose
RUN CGO_ENABLED=0 go build \
    -trimpath \
    -ldflags="-s -w" \
    -o /out/goose \
    github.com/pressly/goose/v3/cmd/goose


# --------------------------------------------------
# Runtime
# --------------------------------------------------

FROM alpine:3.24.1

RUN apk add --no-cache ca-certificates

WORKDIR /app

COPY --from=builder /out/simpleclansstats /app/simpleclansstats
COPY --from=builder /out/goose /usr/local/bin/goose
COPY sql/migrations /app/migrations

# SQLite database
VOLUME ["/data"]

ENV DATABASE_PATH=/data/simpleclansstats.sqlite

COPY docker/entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh

ENTRYPOINT ["/entrypoint.sh"]