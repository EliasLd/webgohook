FROM golang:latest AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o webhook ./cmd/webhook

FROM debian:bookworm-slim

WORKDIR /app

RUN apt-get update && apt-get install -y ca-certificates

COPY --from=build /app/webhook /app/webhook
COPY config/services.json /app/config/services.json

EXPOSE 8081

ENTRYPOINT ["/app/webhook"]
