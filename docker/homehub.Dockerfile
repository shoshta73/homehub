FROM node:22.10.0-alpine3.20 AS frontend-base

FROM frontend-base AS frontend-deps

WORKDIR /app

COPY  ./frontend/package.json ./frontend/yarn.lock* ./frontend/package-lock.json* ./frontend/pnpm-lock.yaml* ./

RUN \
  if [ -f yarn.lock ]; then yarn --frozen-lockfile; \
  elif [ -f package-lock.json ]; then npm ci; \
  elif [ -f pnpm-lock.yaml ]; then corepack enable pnpm && pnpm i --frozen-lockfile; \
  else echo "Lockfile not found." && exit 1; \
  fi

FROM frontend-base AS frontend-builder

WORKDIR /app

COPY ./frontend/vite.config.ts ./vite.config.ts
COPY ./frontend/tsconfig.json ./tsconfig.json
COPY ./frontend/tsconfig.app.json ./tsconfig.app.json
COPY ./frontend/tsconfig.node.json ./tsconfig.node.json
COPY ./frontend/postcss.config.js ./postcss.config.js
COPY ./frontend/tailwind.config.js ./tailwind.config.js
COPY ./frontend/components.json ./components.json

COPY ./frontend/index.html ./index.html
COPY  ./frontend/package.json ./frontend/yarn.lock* ./frontend/package-lock.json* ./frontend/pnpm-lock.yaml* ./
COPY --from=frontend-deps /app/node_modules ./node_modules

COPY ./frontend/public ./public
COPY ./frontend/src ./src

RUN \
  if [ -f yarn.lock ]; then yarn run build; \
  elif [ -f package-lock.json ]; then npm run build; \
  elif [ -f pnpm-lock.yaml ]; then corepack enable pnpm && pnpm run build; \
  else echo "Lockfile not found." && exit 1; \
  fi

FROM golang:1.23.2-alpine3.20 AS backend-builder

WORKDIR /app

ENV CGO_ENABLED=1

RUN apk add --no-cache gcc
RUN apk add --no-cache musl
RUN apk add --no-cache musl-dev

COPY go.mod go.sum ./
RUN go mod download

COPY ./log ./log
COPY ./auth ./auth
COPY ./models ./models
COPY ./exec ./exec

RUN go mod tidy

RUN go build -v -race -o ./bin/ ./...

FROM alpine:3.20

WORKDIR /app

COPY --from=frontend-builder /app/dist ./dist
COPY --from=backend-builder /app/bin/server ./server

EXPOSE 443
EXPOSE 80

CMD ["./server"]
