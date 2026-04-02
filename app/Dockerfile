FROM golang:1.22-alpine AS builder

WORKDIR /src

COPY go.mod ./
COPY internal ./internal
COPY main.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /bin/demo-service .

FROM alpine:3.20

RUN addgroup -S app && adduser -S app -G app

WORKDIR /app

COPY --from=builder /bin/demo-service /app/demo-service

USER app

EXPOSE 8080

ENV PORT=8080
ENV APP_NAME=demo-service
ENV APP_ENV=dev
ENV APP_VERSION=local

ENTRYPOINT ["/app/demo-service"]

