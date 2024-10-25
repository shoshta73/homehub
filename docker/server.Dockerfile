FROM golang:1.23.2-alpine3.20 AS builder

WORKDIR /app

ENV CGO_ENABLED=1

RUN apk add --no-cache gcc
RUN apk add --no-cache musl
RUN apk add --no-cache musl-dev

COPY go.mod go.sum ./
RUN go mod download

COPY ./log ./log
COPY ./exec/server/main.go ./

RUN go mod tidy

RUN go build -v -race -o /app/server .

FROM alpine:3.20

WORKDIR /app

COPY --from=builder /app/server ./server

EXPOSE 3000

CMD ["./server"]
