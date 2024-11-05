FROM golang:1.23.2-alpine3.20 AS build

WORKDIR /app

RUN apk add --no-cache git

ENV CGO_ENABLED=1
RUN apk add --no-cache gcc
RUN apk add --no-cache musl
RUN apk add --no-cache musl-dev

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY ./cmd ./cmd
COPY ./internal ./internal

RUN go build -v -o ./bin/ ./...

FROM alpine:3.20

WORKDIR /app

COPY --from=build /app/bin/api ./bin/server

EXPOSE 3000
EXPOSE 3001

CMD ["./bin/server"]
