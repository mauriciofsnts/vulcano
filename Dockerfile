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
FROM debian:bullseye-slim

RUN addgroup -S bot && adduser -S bot -G bot

USER bot

WORKDIR /app
COPY --from=builder /app/assets /app/assets
COPY --from=builder /app/bot /app/bot

ENTRYPOINT ["/app/bot"]
