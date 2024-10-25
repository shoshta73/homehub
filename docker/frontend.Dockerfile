FROM node:22.10.0-alpine3.20 AS base

FROM base AS deps

WORKDIR /app

COPY  ./frontend/package.json ./frontend/yarn.lock* ./frontend/package-lock.json* ./frontend/pnpm-lock.yaml* ./

RUN \
  if [ -f yarn.lock ]; then yarn --frozen-lockfile; \
  elif [ -f package-lock.json ]; then npm ci; \
  elif [ -f pnpm-lock.yaml ]; then corepack enable pnpm && pnpm i --frozen-lockfile; \
  else echo "Lockfile not found." && exit 1; \
  fi

FROM base AS builder

WORKDIR /app

COPY ./frontend/vite.config.ts ./vite.config.ts
COPY ./frontend/tsconfig.json ./tsconfig.json
COPY ./frontend/tsconfig.app.json ./tsconfig.app.json
COPY ./frontend/tsconfig.node.json ./tsconfig.node.json
COPY ./frontend/postcss.config.js ./postcss.config.js
COPY ./frontend/tailwind.config.js ./tailwind.config.js
COPY ./frontend/components.json ./components.json

COPY ./frontend/index.html ./index.html
COPY ./frontend/public ./public

COPY ./frontend/src ./src

COPY  ./frontend/package.json ./frontend/yarn.lock* ./frontend/package-lock.json* ./frontend/pnpm-lock.yaml* ./
COPY --from=deps /app/node_modules ./node_modules

RUN \
  if [ -f yarn.lock ]; then yarn run build; \
  elif [ -f package-lock.json ]; then npm run build; \
  elif [ -f pnpm-lock.yaml ]; then corepack enable pnpm && pnpm run build; \
  else echo "Lockfile not found." && exit 1; \
  fi

FROM nginx:1.27-alpine3.20

RUN rm -rf /usr/share/nginx/html

COPY --from=builder /app/dist /usr/share/nginx/html

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]
