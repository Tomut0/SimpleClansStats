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
    ./cmd

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

RUN apk add --no-cache ca-certificates nodejs npm

WORKDIR /app

COPY --from=builder /out/simpleclansstats /app/simpleclansstats
COPY --from=builder /out/goose /usr/local/bin/goose
COPY sql/migrations /app/sql/migrations
COPY web/package*.json web/.npmrc /app/web/
WORKDIR /app/web
RUN npm ci
COPY web/ /app/web/
ENV API_URL=/api/v1 APP_NAME=SCStats PAGE_SEPARATOR=" - "
RUN npm run build
WORKDIR /app

# SQLite database
VOLUME ["/data"]

ENV DATABASE_PATH=/data/simpleclansstats.sqlite GOOSE_DRIVER=sqlite3 GOOSE_DBSTRING=/data/simpleclansstats.sqlite
ENV SHADCN_PROJECT_DIR=/app/web
ENV APP_THEME_CSS_PATH=/data/theme.css

COPY deploy/entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh /app/simpleclansstats

ENTRYPOINT ["/entrypoint.sh"]
CMD ["--static"]
