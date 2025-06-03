FROM golang:1.24-alpine AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o app cmd/main.go

FROM alpine:latest

WORKDIR /app
RUN apk --no-cache add ca-certificates tzdata

USER 1000
COPY --from=build --chown=1000:1000 /app/app .



CMD [ "sh", "-c", "/app/app"  ]