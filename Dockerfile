# STAGE: BUILD
FROM golang:1.23 AS builder

WORKDIR /app

COPY go.mod .
COPY go.sum .
COPY Makefile .

COPY assets assets
COPY internal internal
COPY cmd cmd

RUN go mod download

RUN make dist

# STAGE: TARGET
FROM debian:bookworm-slim

RUN apt-get update && \
    apt-get install -y --no-install-recommends ca-certificates && \
    rm -rf /var/lib/apt/lists/*

RUN addgroup --system bot && adduser --system --ingroup bot bot

USER bot

WORKDIR /app
COPY --from=builder /app/assets /app/assets
COPY --from=builder /app/bot /app/bot

ENTRYPOINT ["/app/bot"]
